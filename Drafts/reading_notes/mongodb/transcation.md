# 事务

关注 事务处理的一致性和安全性问题。

### WriteConcern

决定一个写操作落到多少个节点才算成功。 可以防止宕机的时候数据丢失。

0：不关心是否成功

1- 集群最大节点数：写操作需要被复制到指定节点数才算成功。

Majority: 写操作需要被复制到大多数节点（超过一半），才算成功。（`超过一半就算安全`）

all: 所有节点都写到了，才返回成功。这是最安全的做法，但是效率低。不要用这个，这样只要一个节点失败，写操作就不行了，不合理。

j: true 写操作落到 journal 日志文件中才算成功（这么做很安全）。 False ：写操作到达内存就算成功。

**建议**： 普通数据 设置 1， 重要数据设置 `majority` 即可。

### ReadConcern

从哪里读？哪个节点?   ------- readPreference

什么样的数据可以读？关注数据隔离性(哪些数据提交了，哪些数据还没有完全提交)。 ------ readConcern

read preference:

- Primary 默认值，只从主节点读数据 （时效性要求高的场景。比如那种刚写入的数据，用这个好，从节点可能还没更新，读不到）
- Primary preferred: 优先主节点，如果主节点很忙或者挂了，读从节点
- Secondary: 只读从节点 （时效性要求低的场景。 用于读取很早以前就写入的数据）
- Secondray preferred: 优先从节点 （用于读取很早以前就写入的数据）
- Nearest: 最近节点（基于 ping 的时间来判断哪个最近，全世界访问的，适合用这种）

Tag: 它可以设置多个 preference 节点。纯 read preference 只能设置一个。

怎么设置：

- 读的时候 options 可以设置。
- 也可以连接的时候全局设置（一般不这么搞）

read concern:

- availiable：

- local
- majority
- linearizable
- Snapshot

**正确实现读写分离：**

使用 write concern majoriry  

read concern majority  + read preference secondary  组合

## 多文档事务

尽量少用。

因为  事务 = 锁，节点协调，额外开销，性能影响。

**事务** 即  ACID Transcation。

A: Atomicity 原子性 -------  单标文档，复制集多表多行， 分片集群多表多行

C:Consistency 一致性 -------- writeconcern, read concern

I：Isolation 隔离性  ---------  read concern

D：Durability 持久性 --------  journal and replication （先写到日志文件，再写到具体数据文件，这样宕机的时候，有日志可以查）

`事务要在60s内完成，不然会被取消`

多文档事务，必须从主节点中读。





