import {MongoClient} from 'mongodb';
const url = 'mongodb://127.0.0.1/acg_server';
const name = 'acg_server';

function getRandomIntInclusive(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min; // 含最大值，含最小值
}

function update() {

    let _db;

    MongoClient.connect(url,(err,db)=>{
        if (err) throw err;
        console.log('数据库已创建');
        _db = db.db(name);
        const mp = new Map();
        _db.collection('acgs').find().forEach( item => {
            // const likes = getRandomIntInclusive(0,1000);

            // _db.collection('acgs').updateOne({_id:item._id},{$set:{'likes':likes}});
            // item.likes = likes;
            // console.log(item._id);
            for (const tg of item.tags) {
                mp.set(tg.id,tg);
            }

        });

        setTimeout(()=>{

            mp.forEach((value,key)=> {
                console.log('key ' + key);
               // _db.collection('tags').insert(value);
            });

        }, 10000);

    });


}

update();
