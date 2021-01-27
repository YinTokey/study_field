'use strict';

const csv = require('csvtojson');
const path = require('path');
const defpath = path.join(__dirname, '../');

function parseCSV() {
  console.log('导出csv测试');

  csv().fromFile(defpath + '/statics/wallpaper_db_acgs.csv').then(json => {
    console.log(json);
  });
}

module.exports = parseCSV;
