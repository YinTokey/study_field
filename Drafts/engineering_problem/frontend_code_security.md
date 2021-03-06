

## 前端敏感信息安全问题
由于前端代码都是暴露出来的，如果需要存储 秘钥之类的东西，概括起来就是`混淆，压缩，加密`。

从另一个角度来说，这样也不安全，任何敏感信息，不管怎么搞，都不适合存服务端。 应该运行时从服务端换取token。当然即使这么做，也是不安全，攻击方可以代替自己请求服务端。所以这个问题，要结合实际情况，采取对自己最合适的策略。


## 具体策略参考
https://juejin.cn/post/6844903861958737927#heading-3 
## 工具
https://github.com/javascript-obfuscator/javascript-obfuscator  这个工具靠谱


## 加密
涉及到多端加密的，可以使用这个，方便统一，减少整体工作量
https://github.com/skavinvarnan/Cross-Platform-AES


