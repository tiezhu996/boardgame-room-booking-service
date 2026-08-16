# boardgame

桌游吧包厢预约核心服务（纯 Go、内存存储），覆盖包厢容量校验与并发预约。

## 目录结构
```
cmd/boardgame/        程序入口
internal/config/      环境配置
internal/model/       包厢与预约模型
internal/store/       内存包厢存储与锁
internal/service/     预约业务（含并发批量）
internal/worker/      预约统计 worker
```

## 运行与测试
```bash
go build ./...
go test ./...
go test -race ./...
```
