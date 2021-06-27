## 可扩展性

### 垂直扩展（Vertical scaling）

就是让机子更好：更多核的CPU, 更大的内存，更快的磁盘（SSD）

缺点：贵，有上限

### 水平扩展（Horizontal scaling）

用更多普通便宜的机子

### 复制

请求通过负载均衡路由到不同的 server application 上，但是每个  server application 对于同一个请求，应该返回同一个结果，即数据需要有一致性。那么每个  server application 不应该存储用户状态相关的信息，比如session。 这些东西怎么存呢？可以都放 redis
