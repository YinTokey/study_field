'use strict';

// 文件模块
const fs = require('fs');
// 系统路径模块
const path = require('path');
const defPath = path.join(__dirname, '../');

function parseLocalAcg() {
    console.log('导出本地json测试');

    fs.readFile(defPath + '/statics/wallpaper_db_acgs.json', 'utf-8', function(err, data) {
        if (err) {
            console.log(err);
        } else {
            console.log(data);
        }
    });
}

module.exports = parseLocalAcg;
