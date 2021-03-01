import 例子
```
// CommonJS
let { start, exists, readFile } = require('fs')
// 相当于
// let _fs = require('fs')
// let start = _fs.start, exists = _fs.exists, readFile = _fs.readFile

// ES6
import { start, exists, readFile } from 'fs'

```
CommonJS的实质是整体加载fs模块生成一个_fs对象，之后再从对象中分别读取3个方法，称为“运行时加载”。而ES6模块是加载3个方法，称为“编译时加载”




## 参考资料
https://juejin.cn/post/6844904003722018830
《深入浅出node.js》第二章

