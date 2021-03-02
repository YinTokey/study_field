
# ES6 module
CommonJS的实质是整体加载fs模块生成一个_fs对象，之后再从对象中分别读取3个方法，称为“运行时加载”。而ES6模块是加载3个方法，称为“编译时加载”

## export （es6 module）

```javascript
// es6 module
const a = "1"
let b = 2
export {
	a,
  b
}

// 默认输出
export default function(){
  console.log('123')
}
// 相当于
function a(){
  console.log('123')
}
export {a as default}; 
```

## import (es6 module)

```javascript
// es6 module

// 单个引入
import {a , b } from './profile'

// import 是单例模式，下面两个import，对应都是来自同一个module。所以下面这句和上面等价。
import {a } from './profile'
import {b } from './profile'

// 改名引入。 使用的时候，用 x 代替 a
import { a as x } from './profile'


// 整体引入
import * as p form './profile'
// 访问
p.a  ,    p.b 

// 整体加载事基于静态分析的，不能运行时改变
// 下面两行都是不允许的
p.a = 123;
p.c = function(){}

// 复合写法，做模块导出传递(相当于模块继承) （把my_module的模块，借助当前模块，导出到外面，外面引用当前模块即可使用 foo,bar , 不用引用my_module）
export { foo, bar} from 'my_module';

// 默认类型导入
import defaultFn from './module2'
// 相当于
import {default as defaultFn} from './module2'

```

对比commonJS的 require,  import 是编译期执行, 必须在头部写好，不能在块级作用域（比如 具体函数体内）写。 commonJS  的 require 是运行时的，所以可以。import 和编译型语言（Objective-C, swift, C/C++）的 #include, #import 语句差不多一个意思。

## import 动态加载

使用import() 函数 异步加载。算是为了弥补缺乏动态性的缺点。

```javascript
// 下面3 中运行时 import 都可以
async function main() {
  const myModule = await import('./myModule.js');
  const {export1, export2} = await import('./myModule.js');
  const [module1, module2, module3] = 
        await Promise.all([
          import('./module1,js'),
          import('./module2.js'),
          import('./module3.js')
        ])
}
```



## ES6 加载commonJS  模块

使用import 加载 commonJS 时，module.exports属性当做模块的默认输出，即等同于export default。

```javascript
// a.js
module.exports = {
  foo: 'hello',
  bar: 'world'
}

// 在import引入时等同于
export default {
  foo: 'hello',
  bar: 'world'
}


import {readfile} from 'fs' //当'fs'为CommonJS模块时错误

// 整体输入
import * as express from 'express'
const app = express.default();
```



## commonjs 加载 ES6 模块

按照正常姿势使用即可。



## 对比commonjs

ES6 module 比 commonJS  更新，还有一种模块规范，AMD，很少用。

CommonJS模块输出的是一个值的复制，ES6输出的是值的引用。和其他语言一样，值复制就是改变了状态后，再次require使用，还是拿到原始状态，因为是两个东西，内存地址不一样。

commonJS 是一次性加载, 第一次require 会生成一个对象，并缓存起来。再次requrie使用时，直接从缓存中拿，反正值复制嘛，每次拿的都是新的。

es6 module 是动态加载，每次使用都加载，无缓存。



----------------

# Common JS

## exports

`exports` 是上下文的一个对象。 一个文件是一个module, export 是module的一个属性，代表导出的出口。要导出的东西，挂载到 exports 上。

```
module.exports = {
	a,
	b
}
```

模块分类：2大类

Node自身提供的模块，叫`核心模块`。编译时加载的

用户编写的模块，叫`文件模块`。运行时加载的。引入时经历3个步骤：1. 路径分析 2.文件定位 3.编译执行

## 循环加载

```javascript
// a.js
var b = require('b');
// b.js
var a = require('a');
```

避免出现这种，会导致程序跑不了。

ES6 module 的循环 import 是可以跑的，但是可能会有意想不到的结果，所以也要避免。




## 参考资料
https://juejin.cn/post/6844904003722018830
《深入浅出node.js》第二章

