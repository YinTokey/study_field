'use strict';

const dbConfig = require('../../config/db');
const MongoClient = require('mongodb').MongoClient;

module.exports = app => {

    // const grpcList = function(call, callback) {
    //     callback(null, { message: 'random resp a' });
    // };

    //  grpcList(call, callback) {
    //     console.log('request list');
    //     console.log(this.ctx);
    //     const page = call.request.page;
    //     const pageSize = call.request.pageSize;
    //
    //     MongoClient.connect(dbConfig.mongodb.url, function(err, db) {
    //         if (err) throw err;
    //         const dbo = db.db(dbConfig.mongodb.name);
    //         dbo.collection('acgs')
    //             .find({})
    //             .skip(pageSize * page)
    //             .limit(pageSize)
    //             .toArray(function(err, result) { // 返回集合中所有数据
    //                 if (err) throw err;
    //                 console.log('查询结束 : ');
    //                 callback(null, { message: JSON.stringify(result) });
    //                 db.close();
    //             });
    //     });
    //
    //     // callback(null, {message: 'list resp'})
    //
    // }
    //
    // async grpcRandom(call, callback) {
    //     console.log('request random');
    //     const { ctx } = this;
    //     callback(null, { message: 'random resp' });
    //
    // }
};


// module.exports = class GrpcController {
//     // constructor(ctx) {
//     //     console.log('----' + ctx);
//     //     this.ctx = ctx;
//     // }
//
//     async List(call, callback) {
//         console.log('request list');
//         console.log(this.ctx);
//         const page = call.request.page;
//         const pageSize = call.request.pageSize;
//
//         MongoClient.connect(dbConfig.mongodb.url, function(err, db) {
//             if (err) throw err;
//             const dbo = db.db(dbConfig.mongodb.name);
//             dbo.collection('acgs')
//                 .find({})
//                 .skip(pageSize * page)
//                 .limit(pageSize)
//                 .toArray(function(err, result) { // 返回集合中所有数据
//                     if (err) throw err;
//                     console.log('查询结束 : ');
//                     callback(null, { message: JSON.stringify(result) });
//                     db.close();
//                 });
//         });
//
//         // callback(null, {message: 'list resp'})
//
//     }
//
//     async Random(call, callback) {
//         console.log('request random');
//         const { ctx } = this;
//         callback(null, { message: 'random resp' });
//
//     }
//
// };
