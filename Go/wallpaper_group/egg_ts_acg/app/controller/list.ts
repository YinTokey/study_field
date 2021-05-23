import { Controller } from 'egg';

export default class RandomController extends Controller {
    async index() {
        const { ctx, service, config } = this;
        // parseInt(query.limit || config.default_limit, 10)
        const limit = parseInt(ctx.query.limit || config.default_limit, 10);
        const startId = ctx.query.startId;
        const tagId = ctx.query.tagId;

        ctx.logger.info('list 请求 start id: ' + startId + ' limit  ' + limit);

        const result = await service.acg.listData(startId, limit, tagId);

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result.data,
            hasNext: result.hasNext,
            startId: result.startId
        };

    }
}
