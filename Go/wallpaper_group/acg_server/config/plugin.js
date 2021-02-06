'use strict';

/** @type Egg.EggPlugin */
module.exports = {
    // had enabled by egg
    static: {
        enable: true,
    },

    mongoose: {
        enable: true,
        package: 'egg-mongoose',
    },

    routerPlus: {
        enable: true,
        package: 'egg-router-plus',
    },

    grpc: {
        enable: true,
        package: 'egg-grpc',
    },

    grpcServer: {
        enable: true,
        package: 'egg-plugin-grpc-server',
    },

};
