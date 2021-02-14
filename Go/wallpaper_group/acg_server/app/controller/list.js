'use strict';

const Controller = require('egg').Controller;
const _ = require('lodash');
const path = require('path');
const fs = require('fs');
const { v4: uuidv4 } = require('uuid');
const awaitWriteStream = require('await-stream-ready').write;
const sendToWormhole = require('stream-wormhole');

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
        const { ctx, config, service } = this;
        const uid = uuidv4();
        const stream = await ctx.getFileStream();
        const filename = uid + path.extname(stream.filename).toLowerCase();

        ctx.logger.info('upload file ' + stream);

        const folderPath = path.join(__dirname, '../public/upload');

        const target = path.join(folderPath, filename);
        const writeStream = fs.createWriteStream(target);
        try {
            await awaitWriteStream(stream.pipe(writeStream));
            ctx.body = {
                success: true,
                url: filename,
            };
        } catch (err) {
            await sendToWormhole(stream);
            throw err;
        }
    }


}

module.exports = ListController;
