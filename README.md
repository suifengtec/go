# Golang 基础
Go语言基础,测试,相关书籍以及 Web 框架。
## 基础
走马观花的记录。
### 关键词

共有25个关键词,分三类:
```
程序声明(2个)： import, package

程序实体的声明和定义(8个):  type ,struct, interface, var, func,const,chan,map

流程控制(15个):  if, else, for, break,cotinue, switch,case, fallthrough, default, defer,range,goto,return,go,select

```

### 类型系统

基础类型(分布尔，数值，字符串三类19种): 
```
string,bool, byte(uint8的别名),rune(int32的别名), int/uint, unit8/uint8, int16/uint16, int32/uint32, int64/unit64, float32,float64,complex64,copmlex128, uintptr
```

聚合类型(2种):  
```
array, struct
```
引用类型(5种): 
```
pointer,slice,map,func,channel
```

[数组,切片和字典/映射的示例](https://github.com/suifengtec/go/tree/master/array-slice-map)

接口类型(2种): interface,error


### 预置

预置数据类型有 20 种: 基础类型中的数据类型，加上  error;

预置常量有4个:
```
true, false, iota, nil
```

预置函数 13 个 : 
```

new , make, delete,
complex, real,imag,
panic,recover,
len,cap,
append,copy,
close


```

### 操作符

除了赋值,算术,比较,逻辑操作符,地址操作符外,Go 还有类叫接收操作符的操作符, 接收操作符仅有一种符号 `<-`:

```

v1 := <-ch1
v2 := <-ch3
v3 := <-ch2

```


### 其它

表达式,语句,函数,结构体,接口,包,库,工程管理(初始化,格式化,测试,构建)等从略。



## 相关书籍

[点这里地球就会爆炸 :panda_face:,你信吗?](https://github.com/suifengtec/go/blob/master/books.md)


## 示例

### mysql-and-rest-api-and-cli

一个简单的示例程序: 如何通过 CLI 或 REST API 对存储在 MySQL 中的数据进行增删改查。

### goroutine

Go 可让函数或方法并发运行,仅需在方法或函数的前面加上关键词 `go` 即可。

### C 和 Go 语言中的乘法表

同样的功能,类似的代码量,编译为 Windows 上的可执行文件后:

```

C 编译的可执行程序是 53KB;

Go 编译的可执行程序是 1604KB;

```

## Web 框架或工具集




* :+1:[Gin](https://gin-gonic.github.io/gin/) : 是一个非常简约的、精简的框架，它只包含最基本的功能。据说比 Martini 快 39倍。它使用 httprouter 进行请求处理的。

* :thumbsdown:[Martini](https://github.com/olebedev/martini) : 出现的比较早的 Go 语言web框架，在它上面集成第三方支持很容易，Martini 还对路由方法和格式提供了广泛的支持，并支持通配符、变量参数、regex限制等等。Martini 有足够大的群众基础，很多问题容易被解决。但是它最近一次维护是在 2014 年...

* [Web.go](https://github.com/hoisie/web) : 一个非常轻量级的框架。
* [Revel](https://revel.github.io/) : 借鉴了 Play! 这个框架, 它似乎将自己定位为“一站式解决方案”，大部分功能都是预先配置的，并安装了最佳功能。这意味着你可以在没有设置的情况下做很多事情，这对许多创业公司和小团队来说都很有吸引力。
* [Beego](https://beego.me/)： 国人谢孟军做的一个工具化了的全功能 Web 框架，支持模块化开发，在国内很火，它有一个强大的ORM系统，但它在页面缓存（缓存了前几个版本的页面）和多第三方扩展的支持上有所欠缺。
* [Buffalo](https://github.com/gobuffalo/buffalo)： 和Beego, Revel 一样会生成你不熟悉的代码。
* :+1:[Ego](https://github.com/go-ego/ego)：一个基于 Gin 用 Go 编写的全栈 Web 框架，轻量级和高效的前端组件解决方案. 前端编译执行，不影响后端效率。
* :+1::+1::+1:[Iris](https://github.com/kataras/iris)： 比较新,很炫,自称是地球上最快的 Go 语言 Web 框架, [相关文档](https://suifengtec.gitbooks.io/iris-cn/content/)。

