# gidevice-demo

使用 Golang 实现的的与 iOS设备 进行通信的跨平台工具库

## 功能:

* 设备列表查询
* 开发者镜像挂载
* App启/停/卸载/安装
* 端口转发
* 截图
* XCTest
* Syslog
* CrashReport
* 沙盒文件访问
* 定位修改/恢复
* Pair


# Go环境配置

安装Go语言需要配置的环境变量有GOROOT、GOPATH和Path

假设Go安装路径为/usr/local/go

### Windows安装配置go

* (1)配置GOROOT
GOROOT的变量值即为GO的安装目录/usr/local/go
* (2)配置GOPATH
GOPATH的变量值即为存储Go语言项目的路径/usr/local/go/project
* (3)配置Path
Path中有其他安装程序的配置信息，这里再增加一个GO的bin目录/usr/local/go/bin
* (4)验证是否配置成功
```shell
go env
```
(5)通过命令查看安装的Go的版本
```shell
go version
```

### linux安装配置go

1.下载二进制包

```shell
wget https://golang.google.cn/dl/go1.17.6.linux-amd64.tar.gz
```

2.将下载的二进制包解压至 /usr/local目录
```shell
cd /usr/local
tar -zxvf go1.17.6.linux-amd64.tar.gz
```

3.将 /usr/local/go/bin 目录添加至 PATH 环境变量

(1)编辑配置文件
```shell
cd ~
vim /etc/profile
```
(2)加入内容
```shell  
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹
(3)刷新
```shell
source /etc/profile
```


### mac安装配置go

[下载pkg安装包](https://golang.google.cn/dl/)

```shell
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹

```shell
source .bash_profile
```

查看GOROOT环境变量
```shell
echo $GOROOT
```

```shell
go version
```

[homebrew安装方式]
```shell
brew install go
```
### go模块管理

>https://go.dev/doc/tutorial/create-module

```shell
go mod init
```
执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
go.mod 文件来跟踪代码的依赖关系。
该文件仅包含模块的名称和代码支持的 Go 版本。
但是当您添加依赖项时，go.mod 文件将列出您的代码所依赖的版本。
这使构建保持可重复性，并使您可以直接控制要使用的模块版本

mod.go文件内容如下：
```go
module demo

go 1.17

require github.com/360EntSecGroup-Skylar/excelize v1.4.1

require (
	github.com/bunsenapp/go-selenium v0.1.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	github.com/xuri/excelize/v2 v2.5.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/text v0.3.6 // indirect
)

```

[注意]Linux、mac，安装完依赖后引入第三方模块还是会提示，可能会报/usr/local/go/gopath/pkg/mod/ 没有权限

解决办法：
```shell
sudo chmod 777 /usr/local/go/gopath/pkg/mod/
```

## 安装依赖

如果环境已有，直接执行安装依赖

```shell
sudo go mod tidy
```
执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。
执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如:go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行


## 安装wda

iOS设备安装wda并启动

## 安装gidevice-cli

下载通信工具
>https://github.com/electricbubble/gidevice-cli/releases

将工具gidevice放到项目下或配置到系统环境变量

* Devices
```shell
gidevice list
```

* DeveloperDiskImage 四种方式
```shell
$ gidevice mount -l
$ gidevice mount -l -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7
$ gidevice mount -d=/path/.../DeviceSupport/14.4/
$ gidevice mount /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg.signature
```

* Forward 3种方式
```shell
# Default port local=8100 remote=8100
$ gidevice forward
$ gidevice forward -l=9100 -r=9100 -u=设备号
```

## 运行脚本

单元测试框架使用testing
### 运行单元测试
```shell
cd src/Test/
go test -v -run=TestA gdevice_test.go //测试执行TestA 方法
go test -v -run=TestB gdevice_test.go //测试执行TestB 方法
go test -v gdevice_testgdevice_test.go //测试执行unit_test.go文件中的所有方法
go test . // 测试文件夹下所有
```

### 运行main方法
```shell
go run main.go
```

## 生成二进制文件

使用go build 命令来

```shell
go build main.go
```

执行命令后，会在根目录下生成main.sh 运行二进制文件

```shell
./main
```

## 安装vscode go 插件

```shell
go get -v github.com/sqs/goreturns
```

## 依赖包管理

### 查看已安装依赖包列表
```shell
go list -m
```

## 设置国内镜像

### 【linux、mac】

```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
# 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```

### 【windows】
可以在 PowerShell 中设置：：

#### 启用 Go Modules 功能
```
$env:GO111MODULE="on"
```


# 常见问题

* 1.执行go list 提示”go list -m: not using modules“

解决办法：

go依赖模块是通过[go module](https://go.dev/blog/using-go-modules)管理

* 2.执行go env 提示”go: GOPATH entry is relative; must be absolute path: "=/Users/liyinchi/gopath".For more details see: 'go help gopath'“

解决办法：

环境配置文件中gopath和goroot不能一样，在gopath文件夹下新建三个文件夹src、pkg、bin

* 3.执行go get github.com/xuri/excelize/v2 提示”go: could not create module cache: mkdir /usr/local/go/gopath/pkg/mod: permission denied“

 解决办法：

先执行mkdir /usr/local/go/gopath/pkg/mod
再执行sudo go getgithub.com/xuri/excelize/v2

* 4.安装依赖报错在镜像上面没有找到
执行 go get xxxx
go get: github.com/electricbubble/gidevice@v0.6.0: reading https://mirrors.aliyun.com/goproxy/github.com/electricbubble/gidevice/@v/v0.6.0.info: 404 Not Found

解决办法：切换镜像到官方镜像或七牛

```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```


* 5.单元测试testing包


	[源文件]
	以_test 结尾：xxx_test.go

	[测试方法名]
	以Test开头：func TestName(t *testing.T){…}	方法名第一个字母不能是小写字母。
	
	[捕获异常]
	在这些函数中，使用 Error、Fail 或相关方法来发出失败信号。
	要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 TestXxx 函数，如上所述。
	将该文件放在与被测试文件相同的包中。该文件将被排除在正常的程序包之外，但在运行 go test 命令时将被包含。

	有关详细信息，请运行 go help test 和 go help testflag 了解

	[跳过测试]
	可以调用 *T 和 *B 的 Skip 方法，跳过该测试或基准测试
	t.Skip("skipping test in short mode.")



* 6. 查看设备连接

```shell
```
* 7. 

```shell
```
* 8. 

```shell
```


# gidevice-demo

使用 Golang 实现的的与 iOS设备 进行通信的跨平台工具库

## 功能:

* 设备列表查询
* 开发者镜像挂载
* App启/停/卸载/安装
* 端口转发
* 截图
* XCTest
* Syslog
* CrashReport
* 沙盒文件访问
* 定位修改/恢复
* Pair


# Go环境配置

安装Go语言需要配置的环境变量有GOROOT、GOPATH和Path

假设Go安装路径为/usr/local/go

### Windows安装配置go

* (1)配置GOROOT
GOROOT的变量值即为GO的安装目录/usr/local/go
* (2)配置GOPATH
GOPATH的变量值即为存储Go语言项目的路径/usr/local/go/project
* (3)配置Path
Path中有其他安装程序的配置信息，这里再增加一个GO的bin目录/usr/local/go/bin
* (4)验证是否配置成功
```shell
go env
```
(5)通过命令查看安装的Go的版本
```shell
go version
```

### linux安装配置go

1.下载二进制包

```shell
wget https://golang.google.cn/dl/go1.17.6.linux-amd64.tar.gz
```

2.将下载的二进制包解压至 /usr/local目录
```shell
cd /usr/local
tar -zxvf go1.17.6.linux-amd64.tar.gz
```

3.将 /usr/local/go/bin 目录添加至 PATH 环境变量

(1)编辑配置文件
```shell
cd ~
vim /etc/profile
```
(2)加入内容
```shell  
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹
(3)刷新
```shell
source /etc/profile
```


### mac安装配置go

[下载pkg安装包](https://golang.google.cn/dl/)

```shell
# goland
export GOROOT=/usr/local/go
export GOPATH=/usr/local/go/gopath/
export PATH=$PATH:$GOROOT/bin
```
>在gopath文件夹下新建src、bin、pkg三个文件夹

```shell
source .bash_profile
```

查看GOROOT环境变量
```shell
echo $GOROOT
```

```shell
go version
```

[homebrew安装方式]
```shell
brew install go
```
### go模块管理

>https://go.dev/doc/tutorial/create-module

```shell
go mod init
```
执行命令go mod init在当前目录下生成一个go.mod文件，执行这条命令时，当前目录不能存在go.mod文件。如果之前生成过，要先删除；
go.mod 文件来跟踪代码的依赖关系。
该文件仅包含模块的名称和代码支持的 Go 版本。
但是当您添加依赖项时，go.mod 文件将列出您的代码所依赖的版本。
这使构建保持可重复性，并使您可以直接控制要使用的模块版本

mod.go文件内容如下：
```go
module demo

go 1.17

require github.com/360EntSecGroup-Skylar/excelize v1.4.1

require (
	github.com/bunsenapp/go-selenium v0.1.0 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/richardlehane/mscfb v1.0.3 // indirect
	github.com/richardlehane/msoleps v1.0.1 // indirect
	github.com/xuri/efp v0.0.0-20210322160811-ab561f5b45e3 // indirect
	github.com/xuri/excelize/v2 v2.5.0 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/net v0.0.0-20210726213435-c6fcb2dbf985 // indirect
	golang.org/x/text v0.3.6 // indirect
)

```

[注意]Linux、mac，安装完依赖后引入第三方模块还是会提示，可能会报/usr/local/go/gopath/pkg/mod/ 没有权限

解决办法：
```shell
sudo chmod 777 /usr/local/go/gopath/pkg/mod/
```

## 安装依赖

如果环境已有，直接执行安装依赖

```shell
sudo go mod tidy
```
执行go mod tidy命令，它会添加缺失的模块以及移除不需要的模块。
执行后会生成go.sum文件(模块下载条目)。添加参数-v，例如:go mod tidy -v可以将执行的信息，即删除和添加的包打印到命令行


## 安装wda

iOS设备安装wda并启动

## 安装gidevice-cli

下载通信工具
>https://github.com/electricbubble/gidevice-cli/releases

将工具gidevice放到项目下或配置到系统环境变量

* Devices
```shell
gidevice list
```

* DeveloperDiskImage 四种方式
```shell
$ gidevice mount -l
$ gidevice mount -l -u=39xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx7
$ gidevice mount -d=/path/.../DeviceSupport/14.4/
$ gidevice mount /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg /path/.../DeviceSupport/14.4/DeveloperDiskImage.dmg.signature
```

* Forward 3种方式
```shell
# Default port local=8100 remote=8100
$ gidevice forward
$ gidevice forward -l=9100 -r=9100 -u=设备号
```

## 运行脚本

单元测试框架使用testing
### 运行单元测试
```shell
cd src/Test/
go test -v -run=TestA gdevice_test.go //测试执行TestA 方法
go test -v -run=TestB gdevice_test.go //测试执行TestB 方法
go test -v gdevice_testgdevice_test.go //测试执行unit_test.go文件中的所有方法
go test . // 测试文件夹下所有
```

### 运行main方法
```shell
go run main.go
```

## 生成二进制文件

使用go build 命令来

```shell
go build main.go
```

执行命令后，会在根目录下生成main.sh 运行二进制文件

```shell
./main
```

## 安装vscode go 插件

```shell
go get -v github.com/sqs/goreturns
```

## 依赖包管理

### 查看已安装依赖包列表
```shell
go list -m
```

## 设置国内镜像

### 【linux、mac】

```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
# 阿里云
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```

### 【windows】
可以在 PowerShell 中设置：：

#### 启用 Go Modules 功能
```
$env:GO111MODULE="on"
```


# 常见问题

* 1.执行go list 提示”go list -m: not using modules“

解决办法：

go依赖模块是通过[go module](https://go.dev/blog/using-go-modules)管理

* 2.执行go env 提示”go: GOPATH entry is relative; must be absolute path: "=/Users/liyinchi/gopath".For more details see: 'go help gopath'“

解决办法：

环境配置文件中gopath和goroot不能一样，在gopath文件夹下新建三个文件夹src、pkg、bin

* 3.执行go get github.com/xuri/excelize/v2 提示”go: could not create module cache: mkdir /usr/local/go/gopath/pkg/mod: permission denied“

 解决办法：

先执行mkdir /usr/local/go/gopath/pkg/mod
再执行sudo go getgithub.com/xuri/excelize/v2

* 4.安装依赖报错在镜像上面没有找到
执行 go get xxxx
go get: github.com/electricbubble/gidevice@v0.6.0: reading https://mirrors.aliyun.com/goproxy/github.com/electricbubble/gidevice/@v/v0.6.0.info: 404 Not Found

解决办法：切换镜像到官方镜像或七牛

```shell
# 官方
go env -w  GOPROXY=https://goproxy.io
w GOPROXY=https://mirrors.aliyun.com/goproxy/
# 七牛
go env -w  GOPROXY=https://goproxy.cn
# 确认镜像是否成功
go env | grep GOPROXY
```


* 5.单元测试testing包


	[源文件]
	以_test 结尾：xxx_test.go

	[测试方法名]
	以Test开头：func TestName(t *testing.T){…}	方法名第一个字母不能是小写字母。
	
	[捕获异常]
	在这些函数中，使用 Error、Fail 或相关方法来发出失败信号。
	要编写一个新的测试套件，需要创建一个名称以 _test.go 结尾的文件，该文件包含 TestXxx 函数，如上所述。
	将该文件放在与被测试文件相同的包中。该文件将被排除在正常的程序包之外，但在运行 go test 命令时将被包含。

	有关详细信息，请运行 go help test 和 go help testflag 了解

	[跳过测试]
	可以调用 *T 和 *B 的 Skip 方法，跳过该测试或基准测试
	t.Skip("skipping test in short mode.")



* 6. 查看设备连接

```shell
```
* 7. 

```shell
```
* 8. 

```shell
```


