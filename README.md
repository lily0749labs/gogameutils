# gogameutils

`gogameutils` 是面向 Go 游戏服务的工具库，提供扑克牌算法、行为树、定时任务、TCP 封包、数据库与 Redis 连接等基础能力。

模块路径：`github.com/lily0749labs/gogameutils`

## 安装

```bash
go get github.com/lily0749labs/gogameutils@latest
```

Go 1.25 或更高版本。

## 包概览

| 包 | 用途 |
| --- | --- |
| `accountUtil` | 玩家昵称、短信验证码和初始密码生成 |
| `algoUtil` | 三张牌牌型判断与指定牌型发牌 |
| `pokerUtil` | 扑克牌堆、洗牌、发牌和金花牌型编码 |
| `behavior` | 由 JSON 定义条件、概率与动作节点的行为树 |
| `clockUtil` | 支持单次、重复和分类清理的定时任务 |
| `tcp` / `tcp/io` | TCP 客户端、服务端及长度字段封包，支持 JSON 和 Gob |
| `daoUtil/mysqlUtil` | 通过 GORM 连接 MySQL |
| `daoUtil/pgUtil` | 通过 GORM 连接 PostgreSQL |
| `redisUtil` | 创建 go-redis 客户端 |
| `cfgUtil` | 常用服务配置结构 |
| `curUtil` | 分数与货币字符串转换 |
| `lockUtil` | `sync.RWMutex` 封装 |
| `queue` | 带超时的内存队列 |
| `recover` | panic 恢复辅助 |
| `httpUtil` | JSON HTTP 请求 |
| `valiUtil` | 常用参数与格式校验 |

## 快速开始

```go
package main

import (
	"fmt"

	"github.com/lily0749labs/gogameutils/pokerUtil"
)

func main() {
	game := pokerUtil.GamePoker{}
	game.InitPoker()
	game.ShuffleCards()

	card := game.DealCards()
	value, color := pokerUtil.GetCardValueAndColor(card)
	fmt.Printf("card=%x value=%x color=%x remaining=%d\n", card, value, color, game.GetCardsCount())
}
```

## 开发与验证

```bash
go test ./...
go test -race ./...
go vet ./...
```

MySQL、PostgreSQL 和 Redis 只在调用对应连接方法时需要外部服务；单元测试不依赖它们。

## 许可证

Apache-2.0
