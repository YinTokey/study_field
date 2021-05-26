import { Service } from 'egg';
import { v4 } from 'uuid';
import { imageSize } from 'image-size';
import * as amqp from 'amqplib';

/**
 * Acg Service
 */
export default class Acg extends Service {

    async listData(startId, pageSize, tagId) {

        pageSize = parseInt(pageSize, 10);

        let result;
        let nextId;
        let hasNext;

        if (tagId === undefined) {

            if (startId) {
                result = await this.ctx.model.Acg.find({_id:{$gt:startId}}).sort({createAt:-1}).limit(pageSize);
            } else {
                result = await this.ctx.model.Acg.find({}).limit(pageSize);
            }

            const lastObj = result[result.length-1];
            hasNext = result.length === pageSize;
            if (hasNext) {
                nextId = lastObj._id;
            }
            // const key = `list_${page}_${pageSize}`;

            // const cacheResult = await this.ctx.service.cache.get(key);
            // if (!cacheResult) {
            //     result = await this.ctx.model.Acg.find({}, { _id: 0, __v: 0 }).skip(page * pageSize).limit(pageSize).exec();
            //     // 回填 cache
            //     this.ctx.service.cache.set(key, JSON.stringify(result), 3600 * 24);
            // } else {
            //     // redis 数据处理
            //     result = JSON.parse(cacheResult);
            // }


        } else {

            result = await this.ctx.model.Acg.find(
                {'tags.id':tagId},
                { _id: 0, __v: 0 },
                {explain:'executionStats'}
            );
        }

        return {
            data: result,
            startId: nextId,
            hasNext: hasNext
        };
    }


    async random(n, tagId) {
        n = parseInt(n, 10);
        // 随机获取
        let query;

        if (tagId === undefined) {
            query = [{ $sample: { size: n } }];
        } else {
            tagId = parseInt(tagId, 10);

            query = [
                {
                    $match: {
                        tags: {
                            $exists: true
                        }
                    }
                }, {
                    $unwind: {
                        path: '$tags'
                    }
                }, {
                    $match: {
                        'tags.id': tagId
                    }
                }, {
                    $sample: {
                        size: n
                    }
                }
            ];
        }

        const result = await this.ctx.model.Acg.aggregate(query).exec();

        return result.map(acg => {
            // 删除不必要的字段
            delete acg._id;
            delete acg.__v;
            return acg;
        });

    }

    async tags() {
        // 获取tags列表
        const query = [
            {
                $match: {
                    tags: {
                        $exists: true
                    }
                }
            }, {
                $unwind: {
                    path: '$tags'
                }
            }, {
                $group: {
                    _id: '$tags.name',
                    tags: {
                        $push: '$tags'
                    },
                    count: {
                        $sum: 1
                    }
                }
            }
        ];

        const arr = await this.ctx.model.Acg.aggregate(query).exec();

        return arr.map(tag => {
            // 删除不必要的字段
            const first = tag.tags[0];
            tag.name = first.name;
            tag.id = first.id;
            delete tag._id;
            delete tag.tags;
            return tag;
        });
    }

    async clip() {
        try {
            console.log('clip');
            const { config } = this;

            const conn = await amqp.connect(config.rabbitmq,'heartbeat=60');
            const ch = await conn.createChannel();
            const queue = 'acg.ts.task_1';
            const exch = 'acg.ts.exchange';
            const rkey = 'acg.ts.route.key';
            const msg = {
                type:'clip_task',
                data:'image_resize'
            };
            await ch.assertExchange(exch,'direct',{durable:true}).catch(console.error);
            await ch.assertQueue(queue,{durable:true});
            await ch.bindQueue(queue,exch,rkey);
            await ch.publish(exch,rkey,Buffer.from(JSON.stringify(msg)));
            // await ch.sendToQueue(queue,Buffer.from(JSON.stringify(msg)),{persistent:true});
            setTimeout(()=>{
                ch.close();
                conn.close();
            },500);
        } catch (e) {
            console.log('error catch ' + e);
        }
    }

    async query() {
        console.log('query');
        const { config } = this;

        const conn = await amqp.connect(config.rabbitmq,'heartbeat=60');
        const ch = await conn.createChannel();
        const queue = 'acg.ts.task';

        await ch.assertQueue(queue,{durable:true});
        await ch.prefetch(1);
        const msg = await ch.consume(queue,msg => {
            console.log(msg.content.toString());
            ch.ack(msg);
            // ch.cancel('me');

        });
        console.log('here msg',msg);
        setTimeout(()=>{
            ch.close();
            conn.close();
        },500);

    }

    newAcg(url, filePath) {
        const imgSize = imageSize(filePath);

        const acg = new this.ctx.model.Acg();
        acg.imageUrl = url;
        acg.largeImageUrl = url;
        acg.categories = 'acg';
        acg.width = imgSize.width;
        acg.height = imgSize.height;
        acg.pictureId = v4();
        return acg;
    }

    async restoreJSON() {
        // console.log('开始导出本地json测试 1 ');

        // const data = await parseLocalAcg();

        // for (const i in data) {
        //     const obj = data[i];
        //     const po = await this.createPO(obj);
        //     this.ctx.model.Acg.create(po);
        //     console.log(i);
        // }
    }

    // 构建持久化模型
    // async createPO(obj) {
    //     return {
    //         pictureId: obj.picture_id,
    //         imageUrl: obj.image_url,
    //         largeImageUrl: obj.large_image_url,
    //         name: obj.name,
    //         description: obj.description,
    //         author: obj.author,
    //         width: obj.width,
    //         height: obj.height,
    //         likes: obj.likes,
    //         categories: obj.categories,
    //     };
    // }
}
