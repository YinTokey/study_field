'use strict';
const config = require('../../config/db');
const client = require('mongodb').MongoClient;
// const url = 'mongodb://localhost:27017/runoob';
let _db;

client.connect(config.mongodb.url, function(err, db) {
    if (err) throw err;
    console.log('数据库已创建');
    console.info('ni');
    _db = db.db(config.mongodb.name);
});

function mongoClient() {

    return _db;
}

module.exports = mongoClient;
