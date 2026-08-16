# BUG_REPRO

## Bug 是什么
容量判断用不等号导致边界误判；已订数统计错用房间数；批量预约成功计数无锁丢更新；统计总数多 1。

## 如何触发
`go test -race ./...`

## 错误信息
- model.TestCanBook 失败（4,5 被判可订）。
- store.TestReserveRejectsWhenFull 失败（booked=1 期望 2）。
- service.TestBatchReserve 在 `-race` 下报 DATA RACE。
- worker.TestRun 失败（total=2 期望 1）。
