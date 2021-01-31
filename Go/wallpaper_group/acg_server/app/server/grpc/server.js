const grpc = require('@grpc/grpc-js')
const acgProto = require('../../proto/acg_proto')
const config = require('../../../config/grpc')

let cnt = 1

function List(call, callback) {
    console.log("request list")
    callback(null, {message: 'list resp'})
}

function Random(call, callback) {
    console.log("request random")
    callback(null, {message: 'random resp'})
}

function main() {
    const server = new grpc.Server()
    server.addService(acgProto.FetchData.service, {
        List: List,
        Random: Random
    })
    
    server.bindAsync(config.rpc.address + ':' + config.rpc.port, grpc.ServerCredentials.createInsecure(), () => {
        server.start()
        console.log('grpc server started')
        
    })
    
}

main()