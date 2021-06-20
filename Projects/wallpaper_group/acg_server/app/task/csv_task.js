'use strict';

const csv = require('csvtojson');
const path = require('path');
const defPath = path.join(__dirname, '../');

async function parseCSV() {
    console.log('导出csv测试');

    csv().fromFile(defPath + '/statics/wallpaper_db_acgs.csv').then(json => {
        console.log(json);
    });
}

module.exports = parseCSV;
