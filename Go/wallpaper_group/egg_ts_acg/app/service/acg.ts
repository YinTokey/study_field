import { Service } from 'egg';
import { v4 } from 'uuid';
import { imageSize } from 'image-size';

/**
 * Acg Service
 */
export default class Acg extends Service {

    async listData(page, pageSize, tagId) {
        // string to number
        page = parseInt(page, 10);
        pageSize = parseInt(pageSize, 10);

        let query;
        let result;

        if (tagId === undefined) {

            const key = `list_${page}_${pageSize}`;

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

        return result;
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
