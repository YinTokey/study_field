'use strict';
// const grpc = require('@grpc/grpc-js');
// const acgProto = require('./app/proto/acg_proto');
// const config = require('./config/grpc');
// const grpcController = new (require('./app/controller/grpc'))();
const consul = require('./app/server/grpc/consulReg');
const transfer = require('./app/mq/transfer');

module.exports = app => {

    // const server = new grpc.Server();

    app.beforeStart(async () => {

    });

    app.ready(async () => {
        // 启动grpc

        // 微服务注册到consul
        consul.register();

        // 消息队列初始化
        // transfer.initMqConsumer();

    });

    app.beforeClose(async () => {
        //  退出时关闭 grpc server
        app.logger.info('grpc server close before');
        // server.forceShutdown();
    });

    process.on('disconnect', async () => {
        // await app.grpcServer.close();
        app.logger.info('grpc server close');
        // server.forceShutdown();
    });

    process.on('uncaughtException', async err => {
        // await app.grpcServer.close();
        app.logger.error(err);
    });

};

