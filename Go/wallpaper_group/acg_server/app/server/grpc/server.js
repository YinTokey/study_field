const grpc = require('@grpc/grpc-js');
const acgProto = require('../../proto/acg_proto');
const config = require('../../../config/grpc');
const acgImpService = new (require('../../proto/acg_imp'));
//const List = require('../../proto/acg_imp')


function main() {
    const server = new grpc.Server()
    server.addService(acgProto.AcgService.service, {
        List: acgImpService.List,
        Random: acgImpService.Random
    })
    
    server.bindAsync(config.rpc.address + ':' + config.rpc.port, grpc.ServerCredentials.createInsecure(), () => {
        server.start()
        console.log('grpc server started')
        
    })
    
}

main()