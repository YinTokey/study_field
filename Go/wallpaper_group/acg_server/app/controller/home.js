'use strict';

const Controller = require('egg').Controller;

class HomeController extends Controller {
    async index() {
        const {
            ctx,
        } = this;
        ctx.body = 'hi, egg';
        ctx.service.acg.restoreJSON();
    }
}

module.exports = HomeController;
