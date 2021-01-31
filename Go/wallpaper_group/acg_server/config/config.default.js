/* eslint valid-jsdoc: "off" */

'use strict';

/**
 * @param {Egg.EggAppInfo} appInfo app info
 */
module.exports = appInfo => {
    /**
     * built-in config
     * @type {Egg.EggAppConfig}
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
        
    }
    
    return {
        ...config,
        ...userConfig,
    };
};
