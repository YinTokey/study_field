import * as asyncRedis from 'async-redis';

const port = 6379;
const host = '127.0.0.1';

const client = asyncRedis.createClient(
    port,
    host
);

export function redisClient() {

    return client;
}
