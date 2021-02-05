'use strict';

module.exports = app => {

    const httpRouter = app.router;
    const { controller } = app;

    const { home, list, random } = controller;

    httpRouter.get('/', home.index);
    httpRouter.get('/list', list.index);
    httpRouter.get('/random', random.index);

};
