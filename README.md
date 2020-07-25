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
```
go mod管理包 贡献

```

参考文献
```
https://juejin.im/post/5ea186b3e51d45470e2bf88d （go mod基础命令）
https://www.cnblogs.com/sunsky303/p/10710637.html (pkg下面包结构及作用)
```
