'use strict';
const grpc = require('@grpc/grpc-js');
const acgProto = require('../proto/acg_proto');
const dbClient = require('mongodb').MongoClient;

let _conn = null;
let _app = null;

module.exports = app => {
    const { config } = app;

    _app = app;
    const server = new grpc.Server();

    // 连接mongodb
    dataInit(app);

    // 启动grpc
    server.addService(acgProto.AcgService.service, {
        List,
        Random,
    });

    server.bindAsync(config.grpc.endPoint, grpc.ServerCredentials.createInsecure(), () => {
        server.start();
        app.logger.info('grpc server started');
        // app.logger.info(controller);

    });
};

async function dataInit(app) {
    const { config } = app;

    try {
        _conn = await dbClient.connect(config.mongodb.url);
        app.logger.info('mongodb 数据库已连接');
        // const col = conn.db(config.mongodb.name).collection('acg_server');

    } catch (err) {
        app.logger.info('错误：' + err.message);
    } finally {
        // if (conn != null) conn.close();
    }
}

async function List(call, callback) {
    const { config } = _app;

    _app.logger.info('===> list request');
    let page = call.request.page;
    let pageSize = call.request.pageSize;

    if (page === undefined) {
        page = 0;
    }

    if (pageSize === undefined) {
        pageSize = 10;
    }

    const col = _conn.db(config.mongodb.name).collection('acgs');
    // 查询时不返回 '_id' ，'__V' 字段
    const result = await col.find({}, { _id: 0, __v: 0 }).skip(page * pageSize).limit(pageSize).toArray();

    callback(null, { message: '请求成功', data: result });
}

async function Random(call, callback) {
    const { config } = _app;

    _app.logger.info('===> random request');
    let count = call.request.count;

    if (count === undefined) {
        count = 1;
    }

    const col = _conn.db(config.mongodb.name).collection('acgs');
    // 查询时不返回 '_id' ，'__V' 字段

    const result = await col.aggregate([{ $sample: { size: count } }]).toArray();

    callback(null, { message: '请求成功', data: result });
}

