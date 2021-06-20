import { Controller } from 'egg';

export default class HomeController extends Controller {
    public async index() {
        const { ctx } = this;
        ctx.status = 200;
        ctx.body = {
            success: true,
            data: 'home page'
        };
    }
}
