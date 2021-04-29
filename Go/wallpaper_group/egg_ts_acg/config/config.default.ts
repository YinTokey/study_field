import { EggAppConfig, EggAppInfo, PowerPartial } from 'egg';

export default (appInfo: EggAppInfo) => {
    const config = {} as PowerPartial<EggAppConfig>;

    /*
     * override config from framework / plugin
     * use for cookie sign key, should change to your own and keep security
     */
    config.keys = appInfo.name + '_1619527448210_4424';

    // add your egg config in here
    config.middleware = [];

    config.mongoose = {
        url: process.env.EGG_MONGODB_URL || 'mongodb://127.0.0.1/acg_server',
        options: {},
    };

    config.mongodb = {
        url: 'mongodb://127.0.0.1/acg_server',
        name: 'acg_server',
    };

    config.redis = {
        client: {
            host: process.env.EGG_REDIS_HOST || '127.0.0.1',
            port: process.env.EGG_REDIS_PORT || 6379,
            password: process.env.EGG_REDIS_PASSWORD || '',
            db: process.env.EGG_REDIS_DB || '0',
        },
    };

    config.default_page = 1;
    config.default_limit = 20;

    // add your special config in here
    const bizConfig = {
        sourceUrl: `https://github.com/eggjs/examples/tree/master/${appInfo.name}`,
    };

    // the return config will combines to EggAppConfig
    return {
        ...config,
        ...bizConfig,
    };
};
