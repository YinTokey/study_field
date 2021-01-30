'use strict';

const Controller = require('egg').Controller;

class RandomController extends Controller {
    async index() {
        const { ctx } = this;
        const count = ctx.queries.count;

        const data = await ctx.service.acg.random(count);

        ctx.body = data;
    }
}

module.exports = RandomController;
