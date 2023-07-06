# hertz-starter-kit

配置 GOPATH（否则后续 `hz` 命令不可用）

```bash
# .zshrc
export GOPATH=/Users/username/go
export PATH=$PATH:$GOPATH/bin
```

安装 `hz`：

```bash
$ go install github.com/cloudwego/hertz/cmd/hz@latest
```

检查 `hz` 是否安装成功：

```bash
$ hz -v

# hz version v0.6.4
```

如何基于 IDL 生成 Go 项目模板，先创建一个 IDL 文件：

```thrift
// idl/hello.thrift
namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name"); // Add api annotations for easier parameter binding
}

struct HelloResp {
    1: string RespBody;
}


service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}
```

创建项目：

```bash
# Execute under GOPATH
hz new -idl idl/hello.thrift

go mod init

# 注意，需要在 go mod 中 `github.com/apache/thrift` 替换为指定版本
go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0

# Tidy & get dependencies
go mod tidy
```

注意，需要将 `github.com/apache/thrift` 替换为指定版本，否则会导致编译失败（应该是版本不兼容问题，默认会安装最新版本）。也可以手动修改 `go.mod` 文件：

```
replace github.com/apache/thrift => github.com/apache/thrift v0.13.0
```

如果更新了 IDL 文件，只需要运行：

```bash
$ hz update -idl idl/hello.thrift
```

编译项目：

```bash
$ ./build.sh
```

运行项目：

```bash
$ ./output/bootstrap.sh
```

参考：

https://www.cloudwego.io/docs/hertz/tutorials/toolkit/usage/usage-thrift/
