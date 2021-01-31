const grpc = require('@grpc/grpc-js')
const hello_proto = require('./acg_proto')

let cnt = 1

function sayHello(call, callback) {
    console.log("request happend")
    callback(null, {message: `[${cnt++}] echo: ` + call.request.message})
}

function main() {
    var server = new grpc.Server()
    server.addService(hello_proto.Greeter.service, {sayHello: sayHello})
    server.bindAsync('0.0.0.0:50051', grpc.ServerCredentials.createInsecure(), () => {
        server.start()
        console.log('grpc server started')
        
    })
    
}

main()