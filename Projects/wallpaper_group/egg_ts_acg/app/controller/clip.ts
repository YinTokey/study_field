import { Controller } from 'egg';

export default class RandomController extends Controller {
    async index() {
        const { ctx, service } = this;

        await service.acg.clip();

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: 'clip page'
        };
    }

    async query() {

        const { ctx, service } = this;

        await service.acg.query();

        ctx.status = 200;
        ctx.body = {
            success: true,
            data: 'clip query'
        };

    }
}
