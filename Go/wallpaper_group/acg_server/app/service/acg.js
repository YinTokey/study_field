'use strict';

const parseLocalAcg = require('../task/local_json_task.js');

const Service = require('egg').Service;

class AcgService extends Service {
    async restoreJSON() {
        const data = await parseLocalAcg();

        console.log(data);

        for (const i in data) {
            const obj = data[i];
            const po = await this.createPO(obj);
            this.ctx.model.Acg.create(po);
        }
    }
    // 构建持久化模型
    async createPO(obj) {
        const po = {
            pictureId: obj.picture_id,
            imageUrl: obj.picture_id,
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

