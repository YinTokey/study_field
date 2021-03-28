'use strict';

module.exports = app => {
    // const httpRouter = app.router.namespace('/');
    //
    const httpRouter = app.router;
    const { controller, middleware } = app;

    const { home, list, random, tag } = controller;

    const pagination = middleware.pagination();

    httpRouter.get('/', home.index);
    httpRouter.get('/list', pagination, list.index);
    httpRouter.get('/random', random.index);
    httpRouter.get('/tags', tag.index);
    httpRouter.post('/upload', list.upload);
};
