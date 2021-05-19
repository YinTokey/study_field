import { Controller } from 'egg';

export default class RandomController extends Controller {
    async index() {
        const { ctx, service, config } = this;
        // parseInt(query.limit || config.default_limit, 10)
        const limit = parseInt(ctx.query.limit || config.default_limit, 10);
        const page = parseInt(ctx.query.page || config.default_page, 10);
        const tagId = ctx.query.tagId;

        ctx.logger.info('list 请求' + page + ' limit  ' + limit);

        const result = await service.acg.listData(page, limit, tagId);

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result
        };

    }
}
