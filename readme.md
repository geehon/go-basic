# Golang 基础学习

## 基础语法


## 项目管理

### 解读 Go Modules 应用

#### go mod 命令的使用
  
    `go mod init`：初始化go mod， 生成go.mod文件，后可接参数指定 module 名。

    `go mod download`：手动触发下载依赖包到本地cache（默认为$GOPATH/pkg/mod目录）

    `go mod graph`： 打印项目的模块依赖结构
