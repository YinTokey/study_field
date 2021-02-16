'use strict';


const Consul = require('consul');

class ConsulReg {

    static register() {
        const consul = new Consul({
            host: '127.0.0.1',
            port: 8500,
            promisify: true,
        });

        console.log('acg.service 开始注册');

        consul.agent.service.register({
            name: 'acg.service',
            address: '127.0.0.1',
            port: 50051,
        }, function(err, result) {
            if (err) {
                console.log('acg.service 注册失败');

                console.error(err);
                throw err;
            }

            console.log('acg.service 注册成功');
        });

    }

}

module.exports = ConsulReg;
