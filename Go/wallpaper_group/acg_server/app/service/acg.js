'use strict';

// const parseLocalAcg = require('../task/local_json_task.js');

const Service = require('egg').Service;

class AcgService extends Service {

    async listData(page, pageSize) {
        // string to number
        const pageNum = parseInt(page, 10);
        const pageSizeNum = parseInt(pageSize, 10);

        const query = {};
        const filter = { _id: 0, __v: 0 };
        const opt = { skip: pageNum * pageSizeNum, limit: pageSizeNum };

        // 查询时不返回 '_id' ，'__V' 字段
        return this.ctx.model.Acg.find(query, filter, opt).exec();

    }


    async random(n) {
        const num = parseInt(n, 10);
        // 随机获取
        const query = [{ $sample: { size: num } }];

        const arr = await this.ctx.model.Acg.aggregate(query).exec();

        return arr.map(acg => {
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
    async createPO(obj) {
        return {
            pictureId: obj.picture_id,
            imageUrl: obj.image_url,
            largeImageUrl: obj.large_image_url,
            name: obj.name,
            description: obj.description,
            author: obj.author,
            width: obj.width,
            height: obj.height,
            likes: obj.likes,
            categories: obj.categories,
        };
    }

}

module.exports = AcgService;
