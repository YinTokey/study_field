'use strict';

/**
 * @param {Egg.Application} app - egg application.js
 */
module.exports = app => {
    require('./router/http')(app);
    require('./router/grpc')(app);

};

