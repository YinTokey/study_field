'use strict';

const Controller = require('egg').Controller;

class RandomController extends Controller {
    async index() {
        const {
            ctx,
        } = this;
        ctx.body = 'hi, random';
    }
}

module.exports = RandomController;
