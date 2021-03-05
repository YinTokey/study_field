



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




