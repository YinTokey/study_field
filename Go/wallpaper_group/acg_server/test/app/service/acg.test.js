'use strict';

const { app, assert } = require('egg-mock/bootstrap');

describe('get()', () => {

    before(() => console.log('order 1'));
    before(() => console.log('order 2'));
    after(() => console.log('order 6'));
    beforeEach(() => console.log('order 3'));
    afterEach(() => console.log('order 5'));

    it('should get exists list', async () => {

        // 创建 ctx
        const ctx = app.mockContext();
        // 通过 ctx 访问到 service.acg

        const page = '1';
        const limit = '20';

        const data = await ctx.service.acg.listData(page, limit);
        assert(data);
    });

    it('should get exists random', async () => {
        const ctx = app.mockContext();

        const count = '1';
        const data = await ctx.service.acg.random(count);

        assert(data);
    });
});
