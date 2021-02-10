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
        // const arr3 = await test.find({ $and: [
        //     { age: { $lt: 17 } },
        //     { name: 'mike' }
        // ] }).toArray();
        // app.logger.info(arr3);

        // 更新数据
        // await test.updateOne({ name: 'mike' }, { $set: { name: 'jack' } });
        // const arr = await test.find().toArray();
        // app.logger.info(arr);
        //
        // // 更新多条数据
        // await test.updateMany({ age: 17 }, { $set: { age: 16 } });
        //
        // // 删除
        // await test.deleteOne({ age: 15 });

        // 排序
        // const arr = await test.find().sort({ name: -1 }).toArray();
        // app.logger.info(arr);

        // 聚合 -- 数字统计
        // const c = await test.countDocuments({ age: 16 });
        // app.logger.info(c);

        //  聚合 -- distinct 去重
        // const arr = await test.distinct('age');
        // app.logger.info(arr);

        // 聚合 -- group
        const arr = await test.aggregate([
            { $match: { age: 16 } }, // 查询条件
            { $group: { _id: '$cust_id', total: { $sum: '$amount' } } }, // 操作
            { $sort: { total: -1 } }
        ]).toArray();
        app.logger.info(arr);

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
