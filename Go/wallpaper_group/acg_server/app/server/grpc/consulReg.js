'use strict';


const Consul = require('consul');

class ConsulReg {

    static register() {
        const consul = new Consul({
            host: '192.168.6.128',
            port: 8500,
            promisify: true,
        });

        console.log('acg.service 开始注册');

        consul.agent.service.register({
            name: 'acg.service',
            address: '192.168.20.193',
            port: 3000,
            check: {
                http: 'http://192.168.20.193:3000/health',
                interval: '10s',
                timeout: '5s',
            }
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
