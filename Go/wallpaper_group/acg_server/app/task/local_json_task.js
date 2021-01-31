'use strict';

// 文件模块
const fs = require('fs').promises;
// 系统路径模块
const path = require('path');
const defPath = path.join(__dirname, '../');


async function parseLocalAcg() {
    console.log('开始导出本地json测试');

    const data = await fs.readFile(defPath + '/statics/wallpaper_db_acgs.json', 'utf-8');

    return JSON.parse(data.toString());
}

module.exports = parseLocalAcg;

