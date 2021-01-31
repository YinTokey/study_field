const path = require('path')

const config = {
    rpc: {
        address: '0.0.0.0',
        port: 50051,
        protoDir: path.join(__dirname, '../proto/')
    }
}

module.exports = config