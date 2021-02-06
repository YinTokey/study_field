'use strict';

const Controller = require('egg').Controller;

class RandomController extends Controller {
    async index() {
        const { ctx, service } = this;

        let count = ctx.queries.count;

        if (count === undefined) {
            count = '1';
        }

        const data = await service.acg.random(count);

        ctx.body = data;
    }
}

module.exports = RandomController;
