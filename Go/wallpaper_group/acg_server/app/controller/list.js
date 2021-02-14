'use strict';

const Controller = require('egg').Controller;

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
