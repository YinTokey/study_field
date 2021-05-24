# 1. 对于Any

**1.1** **{} , object, unknown 区别**

- {}: 包含除 null, undefined 之外的所有值。 现在基本不用 {} 类型
- object: 包含所有非基础类型（bool, number , string 这些是基础类型）

**1.2 减少 any 影响范围**

```jsx
function f1(){
  const x: any = expressionReturningFoo(); // 不建议,后续的x都是any了
  processBar(x)
}

function f2(){
  const x = expressionReturningFoo();
  processBar(x as any) // 建议，只有这里是any
}
```

# 2. 最佳实践

**2.1** 优先使用 js 特性，而不是 ts 特性。 比如枚举，namespace, Parameter Properties

```jsx
// 这就是 Parameter Properties
class Person {
  constructor(public name: string)
}
```

**2.2 使用 Object.entries 遍历对象**

```jsx
interface ABC{
  a:string;
  b:string;
  c:string;
}
function foo(abc:ABC){
  for(const [k,v] of Object.entries(abc)){
    console.log(k,v)
  }
}
```

# 3. 类型设计

优先使用 union of interface而非 interfaces of unions。这个要看原文，好好体会。

```jsx
interface LineLayer {
  type: 'line',
  layout: LineLayout,
  paint: LinePaint
}
// 导入的时候，经常会有这种写法，和点语法差不多等价
const {paint} = layer;
const paint = layer.paint;
```

**_brand**: 用来做标记

# 4. 类型推导

**4.1** 简单类型就不要标记了

```jsx
const a: number = 10; // 不建议
const a = 10 // 可自行推导

const obj: {name: string, age: number} = {name:'yj', age: 20} // 不建议
const obj = { name: 'yj', age: 20} // 自动推导
```

**4.2 对于函数，尽量标明返回类型，不要类型推导，会迷惑调研者**

**4.3 type widening**

使用一个常量初始化一个变量，但没有提供类型标注时，typescript需要为你的变量确定一个类型，这个过程叫 widening （个人觉得和类型推导差别不大）, 如下

```jsx
const mixed = ['x', 1]
```
