# Golang 基础学习

## 基础语法


## 项目管理

### 解读 Go Modules 应用

#### go mod 命令的使用
  
    `go mod init`：初始化go mod， 生成go.mod文件，后可接参数指定 module 名。

    `go mod download`：手动触发下载依赖包到本地cache（默认为$GOPATH/pkg/mod目录）

    `go mod graph`： 打印项目的模块依赖结构

## go项目中引用 GitHub 私有仓库

1. 先在私有仓库发布版本

![](https://cdn.jsdelivr.net/gh/geehon/blogImgBed/img/golang_release_private.png)

2. 设置 `GOPRIVATE`

Go Module 的代理站点默认的 repo 是 https://proxy.golang.org/,direct ，由于我们的模块发布在 github 上，所以我们需要提供一种机制来绕过代理站点的私有仓库，这里可以使用 `GOPRIVATE` 来实现，`GOPRIVATE` 环境变量用来表示不对外公开的模块路径。

在开发环境中我们可以按照如下命令设置`GOPRIVATE`，多个值用逗号隔开。下面设置的是我的账户级别，也可以设置成仓库级别，比如 github.com/geehon/common-module。

```bash
go env -w GOPRIVATE=github.com/geehon
```

3. 在构建过程中传递仓库凭证

![](https://cdn.jsdelivr.net/gh/geehon/blogImgBed/img/golang_doc_git_https.png)
按照上图文档说明在 `~/.gitconfig` 文件添加下面这段内容

```yaml
[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
```

## 测试部分
