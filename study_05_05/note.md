# 记录一些容易出错的地方

1. 安装完成之后要配置GOROOT、PATH环境变量
2. 设置GO111MODULE、GOPROXY变量，命令为`go env -w GO111MODULE=on`，因为go默认使用GOPATH的方式寻找包，如果不设置GOMODULE的方式导包会出错
3. 建立新项目之后要使用`go mod init xxx` 初始化包名为xxx
4. 导包时路径从第3步初始化的包开始写即可（仅限于从子包中向main包导入），无需写绝对地址（写绝对地址会出错）


**`_`空标识符用于语法上需要，而逻辑上不需要的地方。因为golang语法上不允许存在未使用的变量声明，而有时逻辑上可能并不需要这个变量，因此空标识符可以解决这个问题**

5. 短变量通常用于函数内部和if、for等初始化的时候，在包级向量初始化的地方不被允许
6. 在打开文件时，如果在子包中定义的文件名，应该将路径写成从main.go所在的位置到文件所在的位置，而不是子包文件所在的位置相对于文件所在的位置
如：子包路径为：
test(directory)
	main.go
	sonPackage
		fileOpen.go
		data1.txt
		data2.txt
**如果在fileopen.go路径写为（./data1.txt,./data2.txt）将会报错文件找不到，因为在执行时以main.go为根路径，所以文件路径应写为（sonPackage/data1.txt,sonPackage/data2.txt），尽管文件和fileOpen.go在同一目录下**

switch语句不同于C语言中，一个case执行完之后默认有一个break不往下执行，如果非得向下执行可以使用fallthrough击穿条件。并且case条件也不一定是固定的字面量，可以是布尔表达式。switch后不一定非得跟字面量，类似于if和else，如：
```golang
func test(x int) {
	switch {
	case x > 0 :
		xxxxxxxx
	case x == 0 :
		xxxxxxxx
	default :
		xxxxxxxx
	}
}
```

**go语言中的方法和函数不是一种，方法是关联了命名类型的函数**
