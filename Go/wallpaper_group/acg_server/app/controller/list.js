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

            await this.updateFileCache(target, filename);

            ctx.body = {
                success: true,
                url: filename,
            };


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
        const { ctx, config } = this;
        console.log('请求开始 +' + config.upload.requestUrl);
        superagent.post(config.upload.requestUrl)
            .timeout(5000)
            .set('Authorization', config.upload.token)
            .attach('smfile', target, filename)
            .then(res => {
                console.log('请求结束');
                console.log(res);
            });

        // const j = JSON.parse(res);
        ctx.logger.info(res);

        // 存入redis
        const expire = 3600; // 1小时过期
        // await this.service.cache.set(key, filename, expire);
    }

    async updateFileCache(target, filename) {
        const { ctx } = this;
        const hash = await this.fileHash(target);
        const key = keyPrefix + hash;

        const value = await this.service.cache.get(key);
        console.log('value 查询 ' + value);
        if (value !== undefined) {
            console.log('准备删除');
            await promiseFs.unlink(value);
        }

        console.log(target);
        // 更新缓存
        const expire = 3600; // 1小时过期
        await this.service.cache.set(key, String(target), expire);
    }
}

module.exports = ListController;
