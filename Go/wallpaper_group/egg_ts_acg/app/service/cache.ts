import {Service} from 'egg';

export default class Cache extends Service {

    async get(key) {
        // const { logger } = this.app;
        const t = Date.now();
        let data = await this.app.redis.get(key);
        if (!data) return;
        data = JSON.parse(data);
        const duration = (Date.now() - t);
        this.app.logger.debug('Cache', 'get', key, (duration + 'ms'));
        return data;
    }

    async set(key, value, seconds) {
        // const { logger } = this.app;
        const t = Date.now();
        value = JSON.stringify(value);
        await this.app.redis.set(key, value, 'EX', seconds);
        const duration = (Date.now() - t);
        this.app.logger.debug('Cache', 'set', key, (duration + 'ms'));
    }

    async incr(key, seconds) {
        // const { logger } = this.app;
        const t = Date.now();
        const result = await this.app.redis.multi().incr(key).expire(key, seconds)
            .exec();
        const duration = (Date.now() - t);
        this.app.logger.debug('Cache', 'set', key, (duration + 'ms'));
        return result[0][1];
    }
}
