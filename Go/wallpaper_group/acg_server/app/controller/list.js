'use strict';

const Controller = require('egg').Controller;

class ListController extends Controller {
    async index() {
        const { ctx } = this;
        const pageSize = ctx.queries.pageSize;
        const page = ctx.queries.page;

        const data = await ctx.service.acg.listData(page, pageSize);

        ctx.body = data;
    }

    async grpcList(call, callback) {
        callback(null, { message: 'grpcList request' });
        // const { ctx } = this;
        // console.log('grpcList request' + ctx);
        //
        // callback(null, { message: 'grpcList request' });
    }

}

module.exports = ListController;
