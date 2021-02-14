'use strict';

const Controller = require('egg').Controller;

class TagController extends Controller {
    async index() {
        const { ctx, service } = this;

        const result = await service.acg.tags();

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result
        };

    }
}

module.exports = TagController;
