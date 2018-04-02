# go-demo-3

# windows 开发环境配置

## 下载安装包
- https://dl.google.com/go/go1.10.windows-386.msi
- 默认安装

## 安装vscode

## 配置vscode开发环境
- 查看-扩展(X)-搜索go-安装

## 写hello world 程序
- 写的过程中会提示安装插件  选择install all

## go run **.go 即可运行

## hello world
- package 名是main的 会生成一个可执行的二进制文件
- package 名不是main的 会生成.a 静态库文件

## delve调试工具
- go get github.com/derekparker/delve/cmd/dlv

#### 错误
- C:\GOPATH\src\github.com\derekparker\delve\pkg\proc\disasm.go:11:14: undefined: ArchInst
- 不支持32位系统

## 语法
- 没有; 编译器会自动加上;

## launch.json 配置环境
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "remotePath": "",
            "port": 2345,
            "host": "127.0.0.1",
            "program": "${workspaceRoot}",
            "env": {
                "GOPATH": "c://GOPATH",
                "GOROOT": "c://Go"
            },
            "args": []
        }
    ]
}
```
