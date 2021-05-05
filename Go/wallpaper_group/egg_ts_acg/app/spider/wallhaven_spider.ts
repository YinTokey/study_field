import {axiosRequest} from '../util/axios_request';
import * as cheerio from 'cheerio';
import {UniqueID} from 'nodejs-snowflake';
import {mongoClient} from '../database/mongodb';
import {redisClient} from '../database/redis';

const toplistBaseUrl = 'https://wallhaven.cc/toplist?';
const maxPage = 1;

const uid = new UniqueID({
    returnNumber: false
});

class Tag {
    id:string|bigint;

    name:string;
}

class PictureObject {
    detailHtmlUrl: string;

    fileSize:string | null;

    width:string;

    height:string;

    thumbUrl:string;

    fullUrl:string;

    fileName:string;

    tags:Tag[];

    description:string;

    category:string;

    likes:string;
}


const itemList:PictureObject[] = [];

async function fetchList(page:number) {
    const htmlUrl = toplistBaseUrl + 'page=' + page;
    const websiteHtml = await axiosRequest.get(htmlUrl) as string;
    const $ = cheerio.load(websiteHtml);
    $('.thumb img').each((i, v) => {

        const smallUrl = v['attribs']['data-src'] || v['attribs'].src;
        const urlArray = smallUrl.split('/');
        const fileName = urlArray[urlArray.length - 1] as string;
        const fileSize = $('.wall-res').eq(i).html();
        const dataWallpaperId = fileName.substring(0,fileName.length-4);
        const detailUrl = `https://wallhaven.cc/w/${dataWallpaperId}`;

        const picture = new PictureObject();
        picture.tags = [];
        picture.fileSize = fileSize;
        picture.thumbUrl = smallUrl;
        picture.detailHtmlUrl = detailUrl;
        picture.fileName = fileName;

        itemList.push(picture);

    });
}

async function fetchDetail(item:PictureObject) {
    const websiteHtml = await axiosRequest.get(item.detailHtmlUrl) as string;
    const $ = cheerio.load(websiteHtml);

    $('.scrollbox img').each((_, v) => {
        const width = v['attribs']['data-wallpaper-width'];
        const height = v['attribs']['data-wallpaper-height'];
        const fullUrl = v['attribs']['src'];
        const desc = v['attribs']['alt'];

        item.width = width;
        item.height = height;
        item.fullUrl= fullUrl;
        item.description = desc;

    });

    // 读取tag
    const map = new Map();
    $('li').each(function(_, v) {
        const attribsClass = v.parent['attribs']['id'] as string;
        if (attribsClass === 'tags') {

            (async ()=>{
                // const tag = v['children'][0].attribs.title;
                const tagName = v['children'][0]['children'][0].data as string;
                let tagId = await redisClient().get(tagName);
                if (!tagId) {
                    tagId = uid.getUniqueID();
                    redisClient().set(tagName,tagId);
                }
                const tagObj:Tag = {
                    id:tagId,
                    name:tagName,
                };
                map.set(tagName,tagObj);
            })();

        }
    });

    for (const v of map.values()) {
        item.tags.push(v);
    }

    // 获取点赞信息
    $('.sidebar-section dt').each(function(_, v) {
        // 分类
        if(v['children'][0].data === 'Category' && v.next) {
            const category = v.next['children'][0].data;
            item.category = category;
        }

        // 点赞数
        if(v['children'][0].data === 'Favorites' && v.next) {
            if (v.next['children'][0].attribs.title === 'User Favorites') {
                item.likes = v.next['children'][0].children[0].data;
            }

        }
    });

}

async function storeToDB() {
    for (const item of itemList) {
        if (!item.width || !item.height || item.tags.length === 0) {
            continue;
        }

        const acg = {
            pictureId: uid.getUniqueID(),
            name:item.fileName,
            description : item.description,
            width : item.width,
            height : item.height,
            imageURL : item.thumbUrl,
            largeImageUrl : item.fullUrl,
            category : item.category,
            tags : item.tags,
            author:''
        };

        mongoClient().collection('acgs').insertOne(acg, (err, _)=> {
            if (err) throw err;
            // console.log('文档插入成功');
        });
    }
}

async function start() {
    console.log(uid.getUniqueID());
    // 爬取略缩图数据
    let i:number;
    for (i = 1; i <= maxPage; i++) {
        console.log(i);
        await fetchList(i);
    }
    // 爬取详情页数据
    for (const item of itemList) {
        if (item.detailHtmlUrl) {
            try {
                await fetchDetail(item);
            }catch (_) {
            }
        }
    }

    /*
     * 打印结果
     * console.log(itemList);
     */

    // 存到db
    await storeToDB();
}

start();

