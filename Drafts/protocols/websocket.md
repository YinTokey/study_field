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
前面涉及的轮询，每一次请求响应结束后，都会关闭连接。下一次请求，客户端依然需要在请求头中携带大量信息。而 HTTP Streaming 流化技术的实现机制是客户端发送一个请求，服务端发送一个持续更新和保持打开的开放响应。每当服务端有客户端可用信息时，就更新响应，但是连接始终保持打开。通过这种方式来规避频繁请求带来的不必要开销。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1goig5qmrt9j30hq0m8jsc.jpg)

HTTP Stream 存在的**缺点**是和灵活的全双工通信还存在着距离，还是以单向通信为主，服务端可以自由地发数据给客户端，但客户端缺没办法。

## 1.5 WebSocket
WebSocket 是一种**全双工**，**高实时**，**双向**，单套接字**长连接**。由一次HTTP请求可升级为 WebSocket 连接，可重用客户端到服务端，服务端到客户端的同一连接。它和 HTTP 同属于计算机网络七层模型的应用层，基于TCP传输，二者的差异性如下。


特性 |HTTP | WebSocket
----|------- | -------
内容 | MIME 消息 | 文本、二进制消息
传输 | 半双工 | 全双工


# 2. 连接管理
WebSocket 协议的具体运作主要分为三部分：**握手建立连接**，**数据传输**，**关闭连接**。

## 2.1 握手建立
### 2.1.1 HTTP 请求升级

在上方图中可以看到 WebSocket 连接建立的大致流程为一次 HTTP的请求与响应，然后便可建立连接。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1gon7y9o5woj30ni0tg40b.jpg)

初始阶段客户端客户端发送 HTTP Upgrade Request，在 Request Header 中告知服务端将升级为 WebSocket。

![](https://tva1.sinaimg.cn/large/e6c9d24egy1gojpx84wxcj20pa0msaku.jpg)

其中红色字段为必选。

握手建立一开始由客户端发送一个 HTTP 请求，在 Header 中携带信息告知服务端升级为 WebSocket。其中红色和绿色为必选信息。

`Connection` 告知服务端为长连接，且需要升级
`Upgrade` 告知服务端升级为 WebSocket。
`Sec-WebSocket-Version`指定使用的 WebSocket 版本，一般是使用最新的版本，具体有哪些版本，可以参考这里 [WebSocket 版本](https://www.iana.org/assignments/websocket/websocket.xml#version-number)，当前最新版本为 13，所以这里指定该字段的值为 13。

`Sec-WebSocket-Key` 是一个必选字段，它是一个随机数。这个字段主要和后面服务端返回的`Sec-WebSocket-Accept` 配套使用，减少一些恶意连接和意外连接。后面会详细介绍

### 2.1.2 Server response
服务端响应

![](https://tva1.sinaimg.cn/large/e6c9d24egy1gojpx8gjvoj20nw0kytk9.jpg)

`Sec-WebSocket-Accept` 是基于客户端传递的 `Sec-WebSocket-Key` 计算而来。它有一个公开的算法
![](https://tva1.sinaimg.cn/large/008eGmZEgy1gon7fdujfhj30x00godi1.jpg)

首先将 `Sec-WebSocket-Key` 和一个固定的字符串常量 GUID 拼接起来。

拼接后的字符串经过 SHA1 计算，并转成 Base64 编码，即为`Sec-WebSocket-Accept` 值。

算法都是公开的，这种处理方式并不是为了加密以保证数据安全。它主要是为了减少一些意外和恶意连接，更具体的可概括为如下：

1. 前端使用 Ajax 发送请求时，请求头中是无法设置 `Sec-WebSocket-Key` 字段的，这样子可以避免 使用 Ajax 发送HTTP 请求时意外升级为 WebSocket。
2. 客户端通过识别`Sec-WebSocket-Accept`是否计算正确，确保知道服务端已理解 WebSocket 协议(也会有意外，比如服务端单纯计算了`Sec-WebSocket-Accept`，但没有做具体的 WebSocket 相关操作)。

## 2.2 消息传输
**Message消息**：一条消息（message）可由一个或多个帧(Frame)组成，很多时候会将帧和消息混用，因为大部分时候一条消息只使用一个帧。

一个帧的构成如下
![](https://tva1.sinaimg.cn/large/e6c9d24egy1gom2suygcej21aw0myasu.jpg)

`FIN` 表示消息的结尾，FIN = 1 即这条消息结束了。
`RSV1, RSV2, RSV3` 一般为都为0
`opcode`表示操作码，它由 4个 bit 组成，即在16进制中取值范围 0~F。它的对应取值如下
**0**：表示这是一个持续性，一条消息由多个帧组成，这是其中一个帧。
**1**：消息数据类型为文本。（这条消息包含的帧都是文本类型）
**2**：消息数据类型为二进制。（这条消息包含的帧都是二进制类型）
**8**：客户端或服务端向对方发送关闭握手
**9**：客户端或服务端向对方发送 ping。
**A**：客户端或服务端向对方发送 pong。
**B~F**：为保留操作码。

![](https://tva1.sinaimg.cn/large/008eGmZEly1gooewqjv5ij30sy0bcab0.jpg)
上图一条消息由3个帧构成，从左往右依次的规律为：
**左** FIN = 0，opcode = 1，表示这是一个文本帧，消息还没结束
**中** FIN = 0, opcode = 0，这是一个持续帧，消息还没结束
**右** FIN = 1，opcode，消息结束，这是该消息最后一帧了。

**如果一条消息只有一个帧组成，那么 opcode 取值必然大于0，FIN的值固定为 1。**

`Payload len`: 用于表示数据长度，一共7位，即取值范围可以是 0~ 2^7(127) 。
数据长度小于等于 125字节，则使用红色圈出的区域即可，剩余的 Extened payload不使用。
数据长度 126~2^16-1 字节，`Payload len`值为 126。
数据长度 2^16 ~ 2^64-1， `Payload len` 值为 127。

## 2.3 心跳管理
通信双端之间需要通过心跳确保对方还处于连接状态，心跳本质上也是一条消息，它含有一个心跳帧。如果识别心跳帧？基于**opcode**。
`opcode = 9`：这是一个ping，可以和普通的帧一样携带数据。
`opcode = A`：这是一个pong (当一端收到 ping 之后，会回复一个 pong 给对方，且必须与ping数据相同)。


## 2.4 关闭连接

WebSocket 是基于 TCP 的，需要先关闭上层 WebSocket 连接，才会关闭 TCP。

![](https://tva1.sinaimg.cn/large/008eGmZEgy1govyfg9logj30mw0lumya.jpg)

要关闭 WebSocket 连接时，A 端 一个`opcode = 8`的关闭帧发送给对方。关闭帧可以携带数据，说明连接关闭的原因。发送关闭帧后，进入 closing 状态，此时可以接受数据，但是无法发送。B 端收到关闭帧后，会回复一个关闭帧，此时不再接受任何消息。A端收到回复后，进入 closed 状态，此时 WebSocket 彻底关闭。

![](https://tva1.sinaimg.cn/large/008eGmZEly1gou6ove8c1j30pc0ng75n.jpg)

关闭帧的 payload 数据前2个字节可以表示关闭会话的原因。
![](https://tva1.sinaimg.cn/large/008eGmZEly1got1gjsphkj31cy0t24qp.jpg)


# 参考资料
https://tools.ietf.org/html/rfc6455
https://time.geekbang.org/course/detail/175-93596
https://en.wikipedia.org/wiki/WebSocket#References
https://www.iana.org/assignments/websocket/websocket.xml#version-number
https://www.cnblogs.com/chyingp/p/websocket-deep-in.html
http://www.adambarth.com/papers/2011/huang-chen-barth-rescorla-jackson.pdf
http://websocket.org/quantum.html
https://halfrost.com/websocket/
https://sookocheff.com/post/networking/how-do-websockets-work/#:~:text=Closing%20a%20WebSocket%20connection%20%E2%80%94%20The,indicates%20the%20reason%20for%20closing.
《HTML5 WebSocket权威指南》


