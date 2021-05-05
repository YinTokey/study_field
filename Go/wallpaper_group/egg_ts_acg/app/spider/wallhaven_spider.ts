import {axiosRequest} from '../util/axios_request';
import * as cheerio from 'cheerio';

const toplistBaseUrl = 'https://wallhaven.cc/toplist?';
const maxPage = 1;

class PictureObject {
    detailHtmlUrl: string;

    fileSize:string | null;

    width:string;

    height:string;

    thumbUrl:string;

    fullUrl:string;

    fileName:string;

    tags:string[];

    description:string;
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
        picture.fileSize = fileSize;
        picture.thumbUrl = smallUrl;
        picture.detailHtmlUrl = detailUrl;
        picture.fileName = fileName;

        itemList.push(picture);

    });
}

async function fetchDetail(url:string,item:PictureObject) {
    const websiteHtml = await axiosRequest.get(url) as string;
    const $ = cheerio.load(websiteHtml);
    // console.log($);
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
}

async function start() {
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
                await fetchDetail(item.detailHtmlUrl,item);
            }catch (e) {
                console.log(e);
            }
        }
    }
    // 打印结果
    console.log(itemList);
}

start();

