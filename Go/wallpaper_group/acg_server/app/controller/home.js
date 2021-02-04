'use strict';

const Controller = require('egg').Controller;
// const client = require('../database/redis');

class HomeController extends Controller {
    async index() {
        const {
            ctx,
        } = this;
        // 从redis中获取
        // const result = await client().get('hello');

        ctx.body = 'home page';
    }
}

module.exports = HomeController;
