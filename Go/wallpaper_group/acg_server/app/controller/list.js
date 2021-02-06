'use strict';

const Controller = require('egg').Controller;

class ListController extends Controller {
    async index() {
        const { ctx, service } = this;

        const limit = ctx.pagination.limit;
        const page = ctx.pagination.page;

        console.log('list 请求' + page + limit);

        const data = await service.acg.listData(page, limit);

        ctx.body = data;
    }

    // async grpcList() {
    //     console.log('---- 空炮');
    //     // callback(null, { message: 'grpcList request' });
    //     // const { ctx } = this;
    //     // console.log('grpcList request' + ctx);
    //     //
    //     // callback(null, { message: 'grpcList request' });
    // }


}

module.exports = ListController;
