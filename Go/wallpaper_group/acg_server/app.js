'use strict';
const grpc = require('@grpc/grpc-js');
const acgProto = require('./app/proto/acg_proto');
const config = require('./config/grpc');

module.exports = app => {

    const server = new grpc.Server();

    app.beforeStart(async () => {
        // 启动grpc
        server.addService(acgProto.AcgService.service, {
            // List: listImp.List,
            // Random: acgImpService.Random,
        });

        server.bindAsync(config.rpc.address + ':' + config.rpc.port, grpc.ServerCredentials.createInsecure(), () => {
            server.start();
            app.logger.info('grpc server started');

        });
    });

    app.beforeClose(async () => {
        //  退出时关闭 grpc server
        app.logger.info('grpc server close before');
        server.close();
    });

    process.on('disconnect', async () => {
        // await app.grpcServer.close();
        app.logger.info('grpc server close');
        server.close();
    });

    process.on('uncaughtException', async err => {
        // await app.grpcServer.close();
        app.logger.error(err);
    });
};

