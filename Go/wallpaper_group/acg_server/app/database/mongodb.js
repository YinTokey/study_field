const config = require('../../config/db')
const client = require('mongodb').MongoClient;
//const url = 'mongodb://localhost:27017/runoob';
let db;

client.connect(config.mongodb.url, function (err, db) {
    if (err) throw err;
    console.log('数据库已创建');
    db = db.db(config.mongodb.name);
});

function mongoClient() {
    
    return db;
}

module.exports = mongoClient;