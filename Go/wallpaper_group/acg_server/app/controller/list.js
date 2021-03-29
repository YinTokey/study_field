'use strict';

const Controller = require('egg').Controller;
const _ = require('lodash');
const path = require('path');
const fs = require('fs');
const promiseFs = require('fs').promises;
const { v4: uuidv4 } = require('uuid');
const awaitWriteStream = require('await-stream-ready').write;
const superagent = require('superagent');

const sendToWormhole = require('stream-wormhole');
const crypto = require('crypto');

const keyPrefix = 'acg_server_uploads_';

class ListController extends Controller {
    async index() {
        const { ctx, service } = this;

        const limit = ctx.query.limit > 0 ? ctx.query.limit : ctx.pagination.limit;
        const page = ctx.query.page > 0 ? ctx.query.page : ctx.pagination.skip;
        const tagId = ctx.query.tag;

        ctx.logger.info('list 请求' + page + '  ' + limit);

        const result = await service.acg.listData(page, limit, tagId);

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result
        };

    }

    async upload() {
        const { ctx } = this;
        const uid = uuidv4();
        const stream = await ctx.getFileStream();
        const filename = uid + path.extname(stream.filename).toLowerCase();
        ctx.logger.info(filename);
        const folderPath = path.join(__dirname, '../public/upload');
        const target = path.join(folderPath, filename);

        const writeStream = fs.createWriteStream(target);

        try {
            await awaitWriteStream(stream.pipe(writeStream));

            ctx.body = {
                success: true,
                url: filename,
            };

            // 异步更新数据
            await this.updateFileCache(target, filename);

        } catch (err) {
            await sendToWormhole(stream);
            throw err;
        }
    }

    async fileHash(target, callback) {
        const stream = fs.createReadStream(target);
        const hash = crypto.createHash('sha1');
        hash.setEncoding('hex');
        // read all file and pipe it (write it) to the hash object
        stream.pipe(hash);

        const end = new Promise(function(resolve, reject) {
            stream.on('end', () => resolve(hash.read()));
            stream.on('error', reject); // or something like that. might need to close `hash`
        });

        return await end;
    }

    // 上传到图床
    async uploadToCloud(key, target, filename) {
        const { config } = this;

        // 上传图床，写入 redis 和 mongodb
        const expire = 3600 * 24;
        const res = await superagent.post(config.upload.requestUrl)
            .timeout(5000)
            .set('origin', config.upload.requestUrl)
            .set('refer', config.upload.requestUrl)
            .attach('image', target);

        // 存入缓存
        await this.service.cache.set(key, String(res.text), expire);

        // 写入 mongodb
        this.writeToDB(res.text, target);

        // 删除临时文件地址
        await promiseFs.unlink(target);


    }

    async updateFileCache(target, filename) {
        const hash = await this.fileHash(target);
        const key = keyPrefix + hash;

        // 读取缓存
        const value = await this.service.cache.get(key);

        if (value !== undefined) {
            const expire = 3600 * 24; // 24小时过期
            // 延长缓存时间
            this.service.cache.incr(key, expire);
        } else {
            // 上传图床，写入 redis 和 mongodb
            await this.uploadToCloud(key, target, filename);

        }

    }

    writeToDB(url, target) {
        const acg = this.service.acg.newAcg(url, target);
        acg.save();
    }

}

module.exports = ListController;
