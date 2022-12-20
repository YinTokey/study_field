const fs = require('fs');
const Json2csvParser = require('json2csv').Parser;


function start(file) {
    let result = []

    const data = fs.readFileSync(file, 'UTF-8');

    // split the contents by new line
    const lines = data.split(/\r?\n/);

    // 题号
    const numPattern = '\\d\\.';
    var numReg = new RegExp(numPattern); // 表示的 \d

    // 百分率
    const persentPattern = '\\s0\\.'
    var persentReg = new RegExp(persentPattern);

    // 补充 % 符号解析，做兼容

    // print all lines
    lines.forEach((line) => {
        let titleStrIndex = -1
        let titleEndIndex = -1

        let obj = {};
        // num
        // title
        // persent
        // difficulty

        const numResult = numReg.exec(line);
        if (numResult) {
            const { index } = numResult;
            const numStr = line.substr(0,numResult.index+1)
            titleStrIndex = index+2
            obj.num = parseInt(numStr)
            // console.log(numStr);
        }

        if (line.indexOf('%') !== -1) {
            const persentIndex = line.indexOf('%');
            const persentIndexStart = persentIndex-4
            titleEndIndex = persentIndexStart - 4
            const persentStr = line.substr(persentIndexStart,persentIndex)
            let persentNum = parseFloat(persentStr)/100
            const arr = line.split('%')
            let difficulty = arr[arr.length-1]
            difficulty = difficulty.trim()
            obj.persent = persentNum
            obj.difficulty = difficulty

        } else {
            const persentResult = persentReg.exec(line);
            if (persentResult) {
                const { index } = persentResult;
                let persentStr = line.substr(index,line.length-1)
                persentStr = persentStr.trim()
                const arr = persentStr.split('\t')
                persentStr = arr[0]
                const persent = parseFloat(persentStr)
                titleEndIndex = index-2
                // console.log(persent);
                obj.persent = persent
                obj.difficulty = arr[1]
            }
        }


        if (titleEndIndex !== -1 && titleStrIndex !== -1) {
            let title = line.substr(titleStrIndex,titleEndIndex)
            title = title.trim()
            const arr = title.split('\t0')
            title = arr[0]
            obj.title = title

        }
        if (obj.num) {
           // console.log(obj)
            result.push(obj)
        }
    });

    return result
}

function over(input) {
    // 序号 300 以上的，按难度分类，按概率低到高排列。每个公司单独成一个 json 文件
    let easyArr = []
    let mediumArr = []
    let hardArr = []
    for(const item of input) {
        if (item.num <= 300) {
            continue
        }
        if (item.persent < 0.3) {
            continue
        }

        if (item.difficulty === '简单') {
            easyArr.push(item)
        } else if (item.difficulty === '中等') {
            mediumArr.push(item)
        } else  {
            hardArr.push(item)
        }
    }

    // 子数组排序
    easyArr.sort(compareDecend('persent'));
    // console.log(easyArr)

    mediumArr.sort(compareDecend('persent'));
    // console.log(hardArr)

    hardArr.sort(compareDecend('persent'));
    // console.log(hardArr)

    // 百分数格式化
    easyArr.forEach(item=>{
        item.persent = (item.persent * 100).toFixed(0) + '%';
    })
    mediumArr.forEach(item=>{
        item.persent = (item.persent * 100).toFixed(0) + '%';
    })
    hardArr.forEach(item=>{
        item.persent = (item.persent * 100).toFixed(0) + '%';
    })

    console.log(easyArr.length)
    console.log(mediumArr.length)
    console.log(hardArr.length)

    const fields = ['num', 'title', 'persent', 'difficulty'];
    const json2csvParser = new Json2csvParser({ fields });

    const csv = json2csvParser.parse(hardArr);

    fs.writeFile("hard.csv", csv, function(err) {
        if(err) {
            return console.log(err);
        }

        console.log("The file was saved!");
    });



    // 组装最后json
    // let lastJson = {}
    // lastJson.company = 'shopee'
    // lastJson.easy = easyArr
    // lastJson.medium = mediumArr
    // lastJson.hard = hardArr

    // const data = JSON.stringify(lastJson, null, 4);
    //
    // fs.writeFileSync('shopee300+.json', data);

}

function compareDecend(p){ //这是比较函数
    return function(m,n){
        var a = m[p];
        var b = n[p];
        return b-a; //降序
    }
}

// 去重
function unique(problems){
    let result = {};
    let finalResult=[];
    for(let i=0;i<problems.length;i++){
        result[problems[i].num]=problems[i];
        //因为songs[i].name不能重复,达到去重效果,且这里必须知晓"name"或是其他键名
    }
    //console.log(result);{"羽根":{name:"羽根",artist:"air"},"晴天":{name:"晴天",artist:"周杰伦"}}
    //现在result内部都是不重复的对象了，只需要将其键值取出来转为数组即可
    for(item in result){
        finalResult.push(result[item]);
    }
    //console.log(finalResult);[{name:"羽根",artist:"air"},{name:"晴天",artist:"周杰伦"}]
    return finalResult;
}

const files = ['tencent.txt','ali.txt','baidu.txt','bytedance.txt','shopee.txt','meituan.txt']

function read() {
    let all = []
    for(const item of files) {
         all = all.concat(start(item))
    }

    // 去重
    all = unique(all)
    // 取 大于 300 ，分类，排序
    over(all)
}

read()


