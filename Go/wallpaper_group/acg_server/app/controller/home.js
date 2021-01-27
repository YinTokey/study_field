'use strict';

const Controller = require('egg').Controller;
const parserCSV = require('../task/csv_task.js');

class HomeController extends Controller {
  async index() {
    const { ctx } = this;
    ctx.body = 'hi, egg';
    parserCSV();
  }
}

module.exports = HomeController;
