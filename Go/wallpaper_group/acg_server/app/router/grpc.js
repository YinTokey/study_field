'use strict';
const grpc = require('@grpc/grpc-js');
const acgProto = require('../proto/acg_proto');

module.exports = app => {
    const { config } = app;

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

    server.bindAsync(config.grpc.endPoint, grpc.ServerCredentials.createInsecure(), () => {
        server.start();
        app.logger.info('grpc server started');
        // app.logger.info(controller);

    });

};
