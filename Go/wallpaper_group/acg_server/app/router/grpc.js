'use strict';
const grpc = require('@grpc/grpc-js');
const acgProto = require('../proto/acg_proto');
const config = require('../../config/grpc');

module.exports = app => {
    console.log('=== grpc router');
    const { controller } = app;

    const server = new grpc.Server();

    // 启动grpc
    server.addService(acgProto.AcgService.service, {
        List: (call, callback) => {
            app.logger.info('===> list request');
            callback(null, { message: '===> list request' });
        },
        Random: (call, callback) => {
            app.logger.info('=== random');
            callback(null, { message: 'random resp' });
        },
    });

    server.bindAsync(config.rpc.address + ':' + config.rpc.port, grpc.ServerCredentials.createInsecure(), () => {
        server.start();
        app.logger.info('grpc server started');
        app.logger.info(controller);

    });

};
