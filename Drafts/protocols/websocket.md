# 1. 历史演进概览
什么是 WebSocket？它是定义客户端和服务端如何通过 Web 进行通信的一种网络协议，该通信协议于2011年被IETF（互联网工程任务组）定为标准RFC 6455，并由RFC7936补充规范。

在 WebSocket 协议之前，还有几种方案可用于实现即时通信。下面将依次介绍。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goiacxvxoij316008ut9j.jpg)

## 1.1 HTTP1.0/1.1

在HTTP 1.0/1.1 中，客户端和服务端建立通信，需要通过发送请求与等待响应。由于 HTTP 是无状态的的，每一次请求和响应都需要在 header 中携带大量信息，并且由于其半双工特性，同一时间流量只能单向流动，即请求发送出去，需要一直等待响应的到来，才能做下一步的操作，一定程度上造成了通信低效。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goig7kifiaj30ia0memy3.jpg)


## 1.2 轮询
在需要显示股票价格走势等这一类对信息实时性要求较高的场景，如果每次都由用户手动去刷新浏览器获取最新信息，会存在明显滞后性。因此可以设定一个定时器，每个一段时间，就像服务端发送请求获取最新结果，这种方法被称为”轮询“（polling）。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goig5rizq1j30j60lkabt.jpg)

如果能明确知道服务端数据更新的间隔，那么轮询是一个不错的方案。然而大部分时候，是无法预测数据什么时候更新的，每个固定周期都发送请求查询，会产生很多不必要的请求，造成浪费。

## 1.3 长轮询
长轮询（long polling）则是基于轮询的另一种方案，它也被称为 Comet 或者反向 AJAX。服务端收到客户端请求后，会保持请求打开，直到有客户端可用的信息或者超时了，才返回可给客户端。这么处理相对于普通的轮询，**优点**是可以减少不必要的请求。但是当信息更新很频繁时，长轮询相对于普通轮询的优势就不再明显。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goig5r42y2j30nu0liabi.jpg)

## 1.4 HTTP Streaming
前面设计的轮询，每一次请求响应结束后，都会关闭连接。下一次请求，客户端依然需要在请求头中携带大量信息。而 HTTP Streaming 流化技术的实现机制是客户端发送一个请求，服务端发送一个持续更新和保持打开的开放响应。每当服务端有客户端可用信息时，就更新响应，但是连接始终保持打开。通过这种方式来规避频繁请求带来的不必要开销。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goig5qmrt9j30hq0m8jsc.jpg)

HTTP Stream 存在的**缺点**是和灵活的全双工通信还存在着距离，还是以单向通信为主，服务端可以自由地发数据给客户端，但客户端缺没办法。

## 1.5 Websocket
Websocket 是一种**全双工**，**高实时**，**双向**，单套接字**长连接**。由一次HTTP请求可升级为 WebSocket 连接，可重用客户端到服务端，服务端到客户端的同一连接。它和 HTTP 同属于计算机网络七层模型的应用层，二者的差异性如下。


特性 |HTTP | WebSocket
----|------- | -------
内容 | MIME 消息 | 文本、二进制消息
传输 | 半双工 | 全双工


# 2. 连接管理
WebSocket 协议的具体运作主要分为两部分：**握手建立连接**和**数据传输**。

## 2.1 握手建立
### 2.1.1 HTTP 请求升级

在上方图中可以看到 WebSocket 连接建立的大致流程为一次HTTP的请求与响应，然后便可建立连接。在这一次请求与响应中，客户端和服务端如何交换消息以建立连接，可看下图的 Request Header 和 Response Header。

![](https://tva1.sinaimg.cn/large/e6c9d24egy1gojpx84wxcj20pa0msaku.jpg)

![](https://tva1.sinaimg.cn/large/e6c9d24egy1gojpx8gjvoj20nw0kytk9.jpg)

握手建立一开始由客户端发送一个 HTTP 请求，在 Header 中携带信息告知服务端升级为 WebSocket。


`Sec-WebSocket-Key` 是一个必选字段，需要确保它是随机的。这个字段的是做什么用的呢？


### 2.1.2 Server response
服务端响应
```
HTTP/1.1 101 Switching Protocols
Upgrade: websocket
Connection: Upgrade
Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=

```
### 3.1.3 连接成功


## 3.2 消息传输

### 3.2.1 消息和帧

### 3.2.2 帧数据

### 3.3.3 掩码

## 3.3 心跳管理


## 3.4 关闭连接

### 3.2.1 收到关闭帧，进入closing状态
此时可以接受数据，但是无法发送


# 3. 会话管理


# 4. 数据格式




# 参考资料
https://tools.ietf.org/html/rfc6455
HTML5 WebSocket权威指南
极客时间 抓包。。

https://en.wikipedia.org/wiki/WebSocket#References

