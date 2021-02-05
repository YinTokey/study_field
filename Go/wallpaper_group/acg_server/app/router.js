'use strict';

/**
 * @param {Egg.Application} app - egg application.js
 */
module.exports = app => {
    const {
        router,
        controller,
    } = app;
    router.get('/', controller.home.index);
    router.get('/list', controller.list.index);
    router.get('/random', controller.random.index);
};

