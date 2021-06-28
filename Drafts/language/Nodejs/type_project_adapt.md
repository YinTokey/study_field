## Typescript 改造指南

https://cnodejs.org/topic/5954a747ff46b8a921c947af

https://juejin.cn/post/6844904008834875400

## Typescript 项目适配
最常见的问题
```
Unknown file extension ".ts"
```
解决方法参照官方的
https://github.com/TypeStrong/ts-node/issues/935

## 代码格式化
tslint已经不使用了，都统一使用 eslint
主要参考这个，再结合具体需要的规则进行补充
https://ts.xcatliu.com/engineering/lint.html


// https://www.cnblogs.com/ssaylo/p/12806757.html



## lodash 使用心得

## 核心

想要对 **对象**，集合，数组，字符串 做什么操作时，直接查文档，看看有没有支持的方法

```jsx
// 生成 a b 之间的随机数
_.random(a,b)

// 取属性名，没有就给默认值
_.get(object,'属性名(可多层)',default)

// 对象深拷贝
_.cloneDeep

// 遍历（集合）元素，返回（断言函数）第一个返回真值的第一个元素
_.find
```
