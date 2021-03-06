'use strict';

const Controller = require('egg').Controller;

class RandomController extends Controller {
    async index() {
        const { ctx, service } = this;

        const count = ctx.query.count > 0 ? ctx.query.count : 1;
        const tagId = ctx.query.tag;

        const result = await service.acg.random(count, tagId);

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result
        };

    }
}

module.exports = RandomController;
