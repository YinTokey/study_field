'use strict';

// const parseLocalAcg = require('../task/local_json_task.js');

const Service = require('egg').Service;

class AcgService extends Service {

    async listData(page, pageSize, tagId) {
        // string to number
        page = parseInt(page, 10);
        pageSize = parseInt(pageSize, 10);
        tagId = parseInt(tagId, 10);

        let query;
        let result;

        if (tagId === undefined) {
            query = {};
            const opt = { skip: page * pageSize, limit: pageSize };
            result = await this.ctx.model.Acg.find(query, opt).exec();
        } else {
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
                    $skip: page * pageSize
                }, {
                    $limit: pageSize
                }
            ];

            result = await this.ctx.model.Acg.aggregate(query).exec();

        }

        return result.map(acg => {
            // 删除不必要的字段
            delete acg._id;
            delete acg.__v;
            return acg;
        });
    }


    async random(n, tagId) {
        n = parseInt(n, 10);
        tagId = parseInt(tagId, 10);
        // 随机获取
        let query;

        if (tagId === undefined) {
            query = [{ $sample: { size: n } }];
        } else {
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
