# goserverutils

Go 服务器工具库。

模块路径：`gogameutils`

本地开发时，从两个仓库的共同父目录使用 `go.work`：

```bash
go test ./gocommutils/... ./goserverutils/...
```

`gocommutils` 发布新版本后，可在本模块中更新正式依赖：

```bash
go get github.com/lily0749labs/goutils@latest
go mod tidy
```
