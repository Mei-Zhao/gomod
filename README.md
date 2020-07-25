# gomod
go mod use test

go mod init ****，开启gomod模式，不在依赖¥GOPATH/src;检查依赖本地是否有，如果没有下载，
在当前目录初始化一个新的module， 就是说将该目录下的工程文件初始化为一个Go Module.
日志：（hello world测试用例）
```
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}
```
```
go run main.go ( 如果是go.mod中没有的文件会下载默认最新版本release 增加go.mod一行)
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
Hello, world.
```
```
module github.com/Mei-Zhao/gomod

go 1.13

require rsc.io/quote v1.5.2 // indirect indirect 表示这个库是间接引用进来的。
```
在本地有缓存，可共享,在$GOPATH/pkg下
Go1.13开始不再推荐使用GOPATH。意思就是说你可以在任何路径下存放你的Go源码文件, 不用再像以前一样非得放到$GOPATH/src中
目录结构：
```
├── github.com
│   └── Mei-Zhao
│       ├── gitpush
│       └── gomod
└── pkg
    ├── mod
    │   ├── cache
    │   ├── golang.org
    │   └── rsc.io
    └── sumdb
        └── sum.golang.org
```

go get与go mod (download)
```
贡献，当你需要获取的代码包是自己需要编码的包，使用go get 或者git clone;对于你依赖的第三方包（一般一个仓库的方式确定的），比如go mod方式，获取第三方依赖，然后可以本地编码
* 在本地新建github.com/qiniu 文件夹 (不需要在src下)
* git clone  git@github.com:qiniu/qshell.git （go mod模式管理第三方依赖）
* 直接go run(go version 1.13.1) main.go 会下载mod依赖（如果本地已有缓存，可以复用）
```
关于如何使用自定义的包
```、
hello
    |--conf
        |-conf.go
    |-main.go
    |-go.mod
如何导入conf 包呢?
先查看go.mod 中的module 后的定义的module_name
在导入时  直接使用module_name/conf   即可
```

新版go get、 build、 run 等都支持module


参考文献
```
https://juejin.im/post/5ea186b3e51d45470e2bf88d （go mod基础命令）
https://www.cnblogs.com/sunsky303/p/10710637.html (pkg下面包结构及作用)
https://www.cnblogs.com/xiaobaiskill/p/11819071.html（goland设置）
```

问题：
直接删除 $GOPATH/pkg/mode 会报错 permission deny???
