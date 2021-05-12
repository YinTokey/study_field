### Cache



https://zhuanlan.zhihu.com/p/24467558

https://developers.google.com/web/fundamentals/performance/optimizing-content-efficiency/http-caching?hl=zh-cn

主要依靠cache-control字段

`max-age=60`表示在接下来的60s 缓存和重用响应

`no-store` 不缓存

### 1 缓存体系 三部分

### 1.1 缓存存储策略

能否被缓存，被谁缓存。 Public、Private、no-cache、max-age 这四个会缓存。

no-store不缓存。

### 1.2 缓存过期策略

决定是否加载缓存。使用Expires字段

```
Expires：当前客户端时间 + maxAge 。
```

Cache-Control 中指定的缓存过期策略优先级高于 Expires，当它们同时存在的时候，后者会被覆盖掉。

### 1.3 缓存对比策略

Last-Modified字段，浏览器拿这个值和服务端对比，对比成功返回304,服务端提示从本地加载数据。否则返回200，从服务端获取数据

