



#### 1. 限制any的使用
- any相当于放弃了类型检测
- any破坏了自动补全
- any对重构代码不友好
- any掩盖了你的类型设计
- 尽你所能避免any 
#### 2. 优先使用类型声明而非类型断言
避免使用装箱类型(String, Number, Boolean, Symbol, BigInt)

```
const a = new String('ss');
const b: string = a; // String无法赋值给string
const c:String = '123' // string可以赋值给String

```
#### 3.公共函数，尽可能类型标注
就像其他语言typedef一样
```
type Binary = (a:number,b:number) =>number;

```

#### 4.充分利用泛型 和 type 来精简代码
如同 swift
#### 5.数组特性
数组实际上是对象，其keys也是string而非number，Typescript里使用number index signature是为了进行更多的类型检查 即使如下代码x[0]和x[‘0’]的行为在运行时完全一致，但是只有x[0]才能正确的推倒出类型。
```
let a : string[] = []
let x = a[0] // x类型为string
let y = a['0'] // 但是y类型为any

```

# 阿宝哥ts指南补充

## 2. 变量类型

注意几种

- never
- unknowen
- any

所有类型也都可以赋值给 `unknown`。这使得 `unknown` 成为 TypeScript 类型系统的另一种顶级类型

`unknown` 类型只能被赋值给 `any` 类型和 `unknown` 类型本身。其他类型复制给unknown声明的变量，会报错。

## 3. 断言

if (!xxx) 这种判空，同时对 null 和 undefined 起作用，所有不要用 if ( xx != undefined) 这种沙雕判断

## 4. type

这个很容易坑

`typeof` 类型保护只支持两种形式：`typeof v === "typename"` 和 `typeof v !== typename`，`"typename"` 必须是 `"number"`， `"string"`， `"boolean"` 或 `"symbol"`。

## 5 类型别名

这个太重要了

```jsx
type Message = string | string[];

let greet = (message: Message) => {
  // ...
};
```

可空参数声明 (使用 **联合类型** )

```jsx
const sayHello = (name: string | undefined) => {
  /* ... */
};
```

类型辨识 和 类型守卫，遇到的时候看一下即可，一般不主动用这些特性。

## 6 交叉类型

通过 `&` 将类型合并起来。

## 9 对象

有个很有意思的特性应用，就是对象展开再组装。

```jsx
let person = {
  name: "Semlinker",
  gender: "Male",
  address: "Xiamen",
};

// 组装对象
let personWithAge = { ...person, age: 33 };

// 获取除了某些项外的其它项
let { name, ...rest } = person;
```

## 10 接口

接口可以被多次定义，然后它们会自然合并。 不过一般不建议这么搞，影响可读性。

```jsx
interface Point { x: number; }
interface Point { y: number; }

const point: Point = { x: 1, y: 2 };
```

## 11 类型

注意  属性和swift一样有 get set 方法

带 abstract 关键字的抽象类

## 12 泛型

### keyof

获取某类型所有键

### in

遍历枚举类型

### infer

这个有点难懂

### extends

对泛型约束

### Partial

就是将某个类型里的属性全部变成可选型

