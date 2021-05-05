import {MongoClient} from 'mongodb';
const url = 'mongodb://127.0.0.1/acg_server';
const name = 'acg_server';
let _db;

MongoClient.connect(url,(err,db)=>{
    if (err) throw err;
    console.log('数据库已创建');
    _db = db.db(name);
});

export function mongoClient() {

    return _db;
}
