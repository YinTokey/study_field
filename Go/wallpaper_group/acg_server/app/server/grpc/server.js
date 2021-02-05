'use strict';
const grpc = require('@grpc/grpc-js');
const acgProto = require('../../proto/acg_proto');
const config = require('../../../config/grpc');
const acgImpService = new (require('../../controller/grpc'))();
const listImp = require('../../proto/acg_list');
// const List = require('../../proto/acg_imp')


function main() {
    const server = new grpc.Server();
    server.addService(acgProto.AcgService.service, {
        List: listImp.List,
        Random: acgImpService.Random,
    });

    server.bindAsync(config.rpc.address + ':' + config.rpc.port, grpc.ServerCredentials.createInsecure(), () => {
        server.start();
        console.log('grpc server started');

    });
}

main();
