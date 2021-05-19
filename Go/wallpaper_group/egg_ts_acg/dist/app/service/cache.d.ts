import { Service } from 'egg';
export default class Cache extends Service {
    get(key: any): Promise<string>;
    set(key: any, value: any, seconds: any): Promise<void>;
    incr(key: any, seconds: any): Promise<any>;
}
