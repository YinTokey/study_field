'use strict';

const { app } = require('egg-mock/bootstrap');
const config = require('../../../config/db');
const client = require('mongodb').MongoClient;

async function dataOperate() {
    let conn = null;
    try {
        conn = await client.connect(config.mongodb.url);
        app.logger.info('数据库已连接');
        const test = conn.db(config.mongodb.name).collection('yin_collection');
        // 查询
        // const arr = await test.findOne({ name: 'mike3' });
        // app.logger.info(arr);

        // 多个条件查询
        // const arr2 = await test.find({ age: { $lt: 17 } }).toArray();
        // app.logger.info(arr2);

        // 组合条件查询
        const arr3 = await test.find({ $and: [
            { age: { $lt: 17 } },
            { name: 'mike' }
        ] }).toArray();
        app.logger.info(arr3);

    } catch (err) {
        app.logger.info('错误：' + err.message);
    } finally {
        if (conn != null) conn.close();
    }
}

describe('test/app/mongodb/tutorial.test.js', () => {
    it('should assert', () => {
        dataOperate();
    });

});
