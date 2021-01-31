'use strict';

const asyncRedis = require("async-redis");
const client = asyncRedis.createClient();

function redisClient() {

    return client;
}

module.exports = redisClient;