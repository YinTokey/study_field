import { Service } from 'egg';
/**
 * Acg Service
 */
export default class Acg extends Service {
    listData(page: any, pageSize: any, tagId: any): Promise<any>;
    random(n: any, tagId: any): Promise<any>;
    tags(): Promise<any>;
    newAcg(url: any, filePath: any): any;
    restoreJSON(): Promise<void>;
}
