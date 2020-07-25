# gomod
A module is defined by a tree of Go source files with a go.mod file in the tree’s root directory.

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

go get与go mod (download)、git clone 
```
第一种姿势：git clone（获取源代码）、git run (go mod download 获取第三方依赖)
贡献，当你需要获取的代码包是自己需要编码的包，使用go get 或者git clone;对于你依赖的第三方包（一般一个仓库的方式确定的），比如go mod方式，获取第三方依赖，然后可以本地编码
* 在本地新建github.com/qiniu 文件夹 (不需要在src下,不创建也可以，因为对于自定义包import 路径依赖 module的定于)
* git clone  git@github.com:qiniu/qshell.git （go mod模式管理第三方依赖）
* 直接go run(go version 1.13.1) main.go 会下载mod依赖（如果本地已有缓存，可以复用）

第二种姿势：(如果不需要.git信息) go mod download 提前fill pkg cache
* 在一个目录下new mod, go mod init test (出现go.mod信息)
 The "go mod download" command is useful mainly for pre-filling
the local cache or to compute the answers for a Go module proxy.
* go mod download -json github.com/qiniu/qshell/v2@v2.4.2 (可以指定版本，默认最新release版本，如果没有最新commit)

第三种姿势：go get -u  （go version 1.13.1）go1.13 go module 依然不是默认打开的，只有在有go.mod的情况下才会使用go module通过proxyrg拉取
 ✘ zhaomei@zhaomeis-MacBook-Pro  ~/tools  go get -u github.com/qiniu/qshell/v2@v2.4.2
go: cannot use path@version syntax in GOPATH mode
初始化 module和不初始化对比
* 初始化下载放在$GOPATH/pkg/mod 下(包括依赖)

* 不初始化$GOPATH/src（包括依赖，包含.git信息）When checking out a new package, get creates the target directory GOPATH/src/<import-path>.
日志 go get -v -u github.com/qiniu/qshell
github.com/qiniu/qshell (download)
github.com/astaxie/beego (download)
github.com/shiena/ansicolor (download)
github.com/aws/aws-sdk-go (download)

commit 8ab5b07e9818cbd3c69162c56dc4bb1c74a59592 (HEAD -> master, origin/master, origin/HEAD)
因为有go.mod 文件go build依然会取download，可以go.mod移除

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
https://blog.csdn.net/dashuniuniu/article/details/103769186 (go mod的使用解说)
https://studygolang.com/articles/26694 (go module release)
```

问题：
直接删除 $GOPATH/pkg/mode 会报错 permission deny???
