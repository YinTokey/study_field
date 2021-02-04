/* eslint valid-jsdoc: "off" */

'use strict';

/**
 * @param {Egg.EggAppInfo} appInfo app info
 */
module.exports = appInfo => {
    /**
     * built-in config
     * @type {{}}
     **/
    const config = exports = {};

    // use for cookie sign key, should change to your own and keep security
    config.keys = appInfo.name + '_1611669606403_8196';

    // add your middleware config here
    config.middleware = [];

    // add your user config here
    const userConfig = {
        // myAppName: 'egg',
    };

    config.mongoose = {
        url: process.env.EGG_MONGODB_URL || 'mongodb://127.0.0.1/acg_server',
        options: {},
    };

    config.grpc = {
        endPoint: '0.0.0.0:50051',
    };

    config.grpcServer = {
        port: 50051, // grpc监听端口
        host: '127.0.0.1', // 监听地址
        timeOut: 5000, // 超时时间
        protoDir: 'app/proto', // proto文件所在文件夹
        grpcDir: 'app/grpc', // 接口实现所在文件夹
        errorHandle(error) { // 全局统一错误处理
            // TODO
            // this 为ctx，接受error参数
        }
    };

    return {
        ...config,
        ...userConfig,
    };
};
