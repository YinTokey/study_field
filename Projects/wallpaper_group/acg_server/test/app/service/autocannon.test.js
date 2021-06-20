'use strict';

const { app, assert } = require('egg-mock/bootstrap');

describe('get()', () => {

    it('should get exists list', async () => {

        // 创建 ctx
        const ctx = app.mockContext();
        // 通过 ctx 访问到 service.acg

        const page = '1';
        const limit = '20';

        const data = await ctx.service.acg.listData(page, limit);
        assert(data);
    });

});
