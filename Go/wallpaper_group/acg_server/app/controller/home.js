'use strict';

const Controller = require('egg').Controller;
const parseLocalAcg = require('../task/local_json_task.js');

class HomeController extends Controller {
    async index() {
        const {
            ctx,
        } = this;
        ctx.body = 'hi, egg';
        parseLocalAcg();
    }
}

module.exports = HomeController;
