# Golang Test

Golang test code.

## Windows env 

```

echo %GOPATH%
set GOPATH=W:\tools\GoPath;Y:\DevSpace\Go
echo %GOPATH%


```

## Project init

```

mkdir src && mkdir bin && mkdir pkg && cd src && touch main.go && git init && git config --local core.autocrlf false && git add . &&  && sublime_text main.go && git add . git commit -m "init" && cd ..


```


## Basic

### 关键词

共有25个关键词,分三类:
```
程序声明(2个)： import, package

程序实体的声明和定义(8个):  type ,struct, interface, var, func,const,chan,map

流程控制(15个):  if, else, for, break,cotinue, switch,case, fallthrough, default, defer,range,goto,return,go,select

```



### 类型系统


基础类型(18种): string,bool, byte,rune, int/uint, unit8/uint8, int16/uint16, int32/uint32, int64/unit64, float32,float64,complex64,copmlex128

复合类型(8 种): Map, Array,Slice, Struct, Interface,Function,Channel,Pointer


### 操作符

除了赋值,算术,比较,逻辑操作符,地址操作符外,Go 还有类叫接收操作符的操作符, 接收操作符仅有一种符号 `<-`:

```

v1 := <-ch1
v2 := <-ch3
v3 := <-ch2

```


### 其它

表达式,语句等从略。

## goroutine

Go 可让函数或方法并发运行,仅需在方法或函数的前面加上关键词 `go` 即可。

## 示例

### mysql-and-rest-api-and-cli

一个简单的示例程序: 如何通过 CLI 或 REST API 对存储在 MySQL 中的数据进行增删改查。