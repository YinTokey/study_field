#### 类型
##### any
any 可以运行时修改变量类型。
```
var a: any = 'aaa';
a = 7;
```
像这样改变类型是可以的。 如果不加any , 它会因为第一次赋值，被当成 string，第二次数字赋值，就报错。
```
let a = 'aaa';
a = 7; //  报错
```
但是有一种方式，可以不报错，默认推断为 any，如下
```
let something;
something = 'seven';
something = 7;
```
一开始直接声明，但是不赋值。
##### 联合类型
联合类型，表示一个变量又多个类型，不是联合体数据结构。
```
let myFavoriteNumber: string | number;
myFavoriteNumber = 'seven';
myFavoriteNumber = 7;

```
访问属性时，如果没有被推断，那么只能访问共用属性。如果已经被推断，就可以访问被推断的类型的属性。


#### interface
```
interface Person {
    firstName: string; // 属性
    lastName: string;
    distanceFromOrigin(point: Point): number; //函数
    width?: number; // 可选型，加个问号
    readonly y: number; //只读属性，只能刚创建时通过构造函数设置初始值
    
}

// interface 可以继承
interface Student extends Person {
    sideLength: number;
}

// 还可以多继承
interface Square extends Shape, PenStroke {
    sideLength: number;
}

// 实现接口
class Clock implements ClockInterface {
    currentTime: Date;
    constructor(h: number, m: number) { }
}

```

**特别点**：typescritp 的接口可以继承类。
```
class Point {
    x: number;
    y: number;
    constructor(x: number, y: number) {
        this.x = x;
        this.y = y;
    }
}

interface Point3d extends Point {
    z: number;
}

let point3d: Point3d = {x: 1, y: 2, z: 3};

```


#### 函数
```
注意参照和返回值写法，有冒号

function greeter(person: string) :string {
    return "Hello, " + person;
}

// 可选参数
function buildName(firstName: string, lastName?: string) {
    if (lastName)
        return firstName + " " + lastName;
    else
        return firstName;
}

// 剩余参数 ， 这点和Go 很像
function buildName(firstName: string, ...restOfName: string[]) {
  return firstName + " " + restOfName.join(" ");
}

let employeeName = buildName("Joseph", "Samuel", "Lucas", "MacKinzie");

```

可选参数必须放必选参数之后。

=> 和 es6的箭头函数有一些区别，左边表示函数参数，右边表示返回值。
```
let mySum: (x: number, y: number) => number = function (x: number, y: number): number {
    return x + y;
};

```

函数也可以被interface 约束，这点相对于其他语言比较特殊。



#### 类
```
注意构造函数

class Student {
    fullName: string;
    constructor(public firstName, public middleInitial, public lastName) {
        this.fullName = firstName + " " + middleInitial + " " + lastName;
    }
}

继承， 还是 extends
class Dog extends Animal {
    breed: string;
}

```
类 自带 `name` 属性。

`protected` 关键字：子类可以访问，外部不可以访问。

```
class Animal {
  public name;
  protected constructor(name) {
    this.name = name;
  }
}

```
当构造函数被声明 protected，它相当于是一个抽象类，不能直接使用，必须有子类基础它，然后使用子类。

**正统的抽象类**
如下，使用 abstract 关键字
```
abstract class Animal {
  public name;
  public constructor(name) {
    this.name = name;
  }
  public abstract sayHi();
}

```

#### 模块导出
原先这么写 (commonJS)
```
function foo() {
    // ...
}
module.exports = foo;
```
现在这么写 export = xxx; (es6 module)
```
function foo() {
    // ...
}
export = foo;
```
#### 模块导入
```
// 对于原js模块 (commonJS)
import fs = require('fs');

// 等价于 原 require 写法 (ES6 module)
import * as ws from 'ws';


// 导入选定的类，函数，属性等 （module里导出了 A, B , C，import 的时候要同名）
import { A , B , C} from './module'

```

#### 类型转换
和swift 一样
```
let options = {} as Options;

```

#### null, undefine, undeclared 区别

undefine 此处应该有值，但是没有
```
var b;
console.log(b); // undefined

var person = {
  name: "elaine",
};
console.log(person.age); // undefined

```

null 仅是一种值的状态，代表空值(空值也是有值)
```
var a = null;
console.log(a); // null
console.log(typeof a); // object (空值也是有值)
```

undeclared 未声明
```
// 嚴格模式，Uncaught ReferenceError: a is not defined 
"use strict";
a = 10;
console.log(a);

// 非严格模式，也会错

```

#### 常量/变量声明
```
和swift 没什么区别
let decLiteral: number = 6;

// 数组，两种声明都可以
let list: Array<number> = [1, 2, 3];
let list: number[] = [1, 2, 3];


```

#### 枚举
和swift差不多。可以不赋值，局部赋值 ，全赋值
```
enum Color {Red = 1, Green, Blue}
let c: Color = Color.Green;
```

枚举有类型枚举，和常量型枚举。要注意 加 const
https://github.com/microsoft/TypeScript/issues/16671

#### object
object表示非原始类型，也就是除number，string，boolean，symbol，null或undefined之外的类型。

#### for循环
**大坑**
```
// for..of则迭代对象的键对应的值
for (let entry of someArray) {
    console.log(entry); // 1, "string", false
}

// for..in迭代的是对象的 键 的列表
for (let i in list) {
    console.log(i); // "0", "1", "2",
}


map 和 set 遍历也可以这么写
```

#### 模块
任何声明都可以export 导出（变量，函数，类，接口，别名等）
```
export interface StringValidator {
    isAcceptable(s: string): boolean;
    
}
```

#### 闭包
闭包回调给外部，可以参照swift的写法
```
class x {
	onInfoGet?: (info: Info) => Info;

}

```
注意可选型的使用。

#### 类型别名
和其他语言的 typedef 差不多
```
type Name = string;
type NameResolver = () => string;
type NameOrResolver = Name | NameResolver;
function getName(n: NameOrResolver): Name {
    if (typeof n === 'string') {
        return n;
    } else {
        return n();
    }
}

```

规定 EventNames 变量只能取这几种类型。
```
type EventNames = 'click' | 'scroll' | 'mousemove';

```

