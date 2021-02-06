'use strict';

// const parseLocalAcg = require('../task/local_json_task.js');

const Service = require('egg').Service;

class AcgService extends Service {

    async listData(page, pageSize) {
        // string to number
        const pageNum = parseInt(page, 10);

        const pageSizeNum = parseInt(pageSize, 10);
        // 查询时不返回 '_id' ，'__V' 字段
        return this.ctx.model.Acg.find({}, { _id: 0, __v: 0 })
            .skip(pageNum * pageSizeNum)
            .limit(pageSizeNum)
            .exec();
    }


    async random(n) {

        const min = 1;
        const max = await this.ctx.model.Acg.find({}).count().exec();
        const nums = this.randomNums(n, min, max);

        const arr = [];
        for (let i = 0; i < nums.length; i++) {
            arr[i] = await this.ctx.model.Acg.find({}, { _id: 0 }).skip(nums[i]).limit(1)
                .exec();
        }

        return arr;
    }

    async restoreJSON() {
        console.log('开始导出本地json测试 1 ');

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


    randomNums(n, min, max) {
        const arr = [];
        for (let i = 0; i < n; i++) {
            let ran = Math.ceil(Math.random() * (max - min) + min);
            while (this.isExist(arr, ran)) {
                ran = Math.ceil(Math.random() * (max - min) + min);
            }
            arr[i] = ran;
        }
        return arr;
    }


    isExist(arr, ran) {
        for (let i = 0; i < arr.length; i++) {
            // eslint-disable-next-line eqeqeq
            if (arr[i] === ran) {
                return true;
            }
        }
        return false;
    }

}

module.exports = AcgService;
