'use strict';

module.exports = {
    grpcList(call, callback) {
        // this 就是 app 对象，在其中可以调用 app 上的其他方法，或访问属性
        console.log('---- 访问xx in context');
        console.log(this.app);
        callback(null, { message: 'LIST resp a' });

    },
};
