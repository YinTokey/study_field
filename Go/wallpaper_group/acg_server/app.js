'use strict';

module.exports = app => {
    const registryClient = app.cluster(app.RealClient, { isBroadcast: false }).create({});

    app.beforeStart(async () => {
        // worker 订阅启动grpc 服务消息
        registryClient.subscribe({
            dataId: 'grpc.server.leader.attach'
        }, val => {
            app.logger.info(`worker subscribe [grpc.server.leader.attach]: ${val}`);
            if (val === process.pid) {
                const grpcServer = new app.GrpcServer(app);
                grpcServer.start();
                Reflect.defineProperty(app, 'grpcServer', { value: grpcServer });
                // 向leader发送start消息
                registryClient.publish({
                    dataId: 'grpc.server.worker.start',
                    publishData: process.pid
                });
            }
        });

        // 向leader发送ready消息
        registryClient.publish({
            dataId: 'grpc.server.worker.ready',
            publishData: process.pid
        });
    });

    // 退出时关闭 grpc server 并 像 leader 发送 close 信号
    app.beforeClose(async () => {
        await app.grpcServer.close();
        // 向leader发送ready消息
        registryClient.publish({
            dataId: 'grpc.server.worker.close',
            publishData: process.pid
        });
    });

    process.on('disconnect', async () => {
        await app.grpcServer.close();
        registryClient.publish({
            dataId: 'grpc.server.worker.close',
            publishData: process.pid
        });
    });

    process.on('uncaughtException', async err => {
        await app.grpcServer.close();
        registryClient.publish({
            dataId: 'grpc.server.worker.close',
            publishData: process.pid
        });
        app.logger.error(err);
    });
};

// class AppBootHook {
//     constructor(app) {
//         this.app = app;
//     }
//
//     const registryClient = this.app.cluster(this.app.RealClient, { isBroadcast: false }).create({})
//
//     configWillLoad() {
//         // logger.info('wiston logg');
//         // 此时 config 文件已经被读取并合并，但是还并未生效
//         // 这是应用层修改配置的最后时机
//         // 注意：此函数只支持同步调用
//
//         // 例如：参数中的密码是加密的，在此处进行解密
//         // this.app.config.mysql.password = decrypt(this.app.config.mysql.password);
//         // // 例如：插入一个中间件到框架的 coreMiddleware 之间
//         // const statusIdx = this.app.config.coreMiddleware.indexOf('status');
//         // this.app.config.coreMiddleware.splice(statusIdx + 1, 0, 'limit');
//     }
//
//     async didLoad() {
//         // 所有的配置已经加载完毕
//         // 可以用来加载应用自定义的文件，启动自定义的服务
//         // 例如：创建自定义应用的示例
//         // this.app.queue = new Queue(this.app.config.queue);
//         // await this.app.queue.init();
//         //
//         // // 例如：加载自定义的目录
//         // this.app.loader.loadToContext(path.join(__dirname, 'app/tasks'), 'tasks', {
//         //     fieldClass: 'tasksClasses',
//         // });
//
//         // worker 订阅启动grpc 服务消息
//         registryClient.subscribe({
//             dataId: 'grpc.server.leader.attach'
//         }, val => {
//             app.logger.info(`worker subscribe [grpc.server.leader.attach]: ${val}`)
//             if (val === process.pid) {
//                 const grpcServer = new app.GrpcServer(app)
//                 grpcServer.start()
//                 Reflect.defineProperty(app, 'grpcServer', { value: grpcServer })
//                 // 向leader发送start消息
//                 registryClient.publish({
//                     dataId: 'grpc.server.worker.start',
//                     publishData: process.pid
//                 })
//             }
//         })
//
//         // 向leader发送ready消息
//         registryClient.publish({
//             dataId: 'grpc.server.worker.ready',
//             publishData: process.pid
//         })
//     }
//
//     async willReady() {
//         // 所有的插件都已启动完毕，但是应用整体还未 ready
//         // 可以做一些数据初始化等操作，这些操作成功才会启动应用
//
//         // 例如：从数据库加载数据到内存缓存
//         // this.app.cacheData = await this.app.model.query(QUERY_CACHE_SQL);
//     }
//
//     async didReady() {
//         // 应用已经启动完毕
//
//         // const ctx = await this.app.createAnonymousContext();
//         // await ctx.service.Biz.request();
//     }
//
//     async serverDidReady() {
//         // http / https server 已启动，开始接受外部请求
//         // 此时可以从 app.server 拿到 server 的实例
//         //
//         // this.app.server.on('timeout', socket => {
//         //     // handle socket timeout
//         // });
//     }
// }
//
// module.exports = AppBootHook;
