import { Controller } from 'egg';

export default class RandomController extends Controller {
    async index() {
        const { ctx, service } = this;

        const count = parseInt(ctx.query.count, 10) > 0 ? parseInt(ctx.query.count, 10) : 1;
        const tagId = ctx.query.tag;

        const result = await service.acg.random(count, tagId);

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: result
        };

    }
}
