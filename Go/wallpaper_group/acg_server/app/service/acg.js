'use strict';

// const parseLocalAcg = require('../task/local_json_task.js');

const Service = require('egg').Service;

class AcgService extends Service {

    async listData(page, pageSize) {
        // string to number
        const pageNum = parseInt(page, 10);
        const pageSizeNum = parseInt(pageSize, 10);


        const data = this.ctx.model.Acg.find({})
            .skip(pageNum * pageSizeNum)
            .limit(pageSizeNum)
            .exec();
        return data;
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
        const po = {
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
        return po;
    }
}

module.exports = AcgService;