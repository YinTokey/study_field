/* eslint valid-jsdoc: "off" */

'use strict';

const path = require('path');

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
    config.middleware = [ 'errorHandler' ];

    // add your user config here
    const userConfig = {
        // myAppName: 'egg',
    };

    config.redis = {
        client: {
            host: process.env.EGG_REDIS_HOST || '127.0.0.1',
            port: process.env.EGG_REDIS_PORT || 6379,
            password: process.env.EGG_REDIS_PASSWORD || '',
            db: process.env.EGG_REDIS_DB || '0',
        },
    };

    config.mongoose = {
        url: process.env.EGG_MONGODB_URL || 'mongodb://127.0.0.1/acg_server',
        options: {},
    };

    config.mongodb = {
        url: 'mongodb://127.0.0.1/acg_server',
        name: 'acg_server',
    };

    config.grpc = {
        endPoint: '0.0.0.0:50051',
    };

    config.upload = {
        path: path.join(__dirname, '../app/public/upload/'),
        // url: '/public/upload/',
        // 图床上传地址
        requestUrl: 'https://img.vim-cn.com',
        token: process.env.UPLOAD_TOKEN
    };
    config.security = {
        csrf: {
            // 调试模式 关闭 csrf
            enable: false,
        },
    };

    config.default_page = 1;
    config.default_limit = 20;

    return {
        ...config,
        ...userConfig,
    };
};
