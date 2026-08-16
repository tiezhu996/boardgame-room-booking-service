# BUG_REPRO

## Bug 是什么
配置读取返回空；容量判断只检查容量为正导致超订；预约不再校验容量；批量预约 wg.Add 放进 goroutine 且计数无锁丢更新；统计用房间数翻倍。

## 如何触发
`go test -race ./...`

## 错误信息
- config.TestLoad 失败（AppName 为空）。
- model.TestCanBook 失败（超订）。
- service.TestBatchReserve 失败且 `-race` 报 DATA RACE。
- store.TestReserveRejectsWhenFull 失败（满员仍预约）。
- worker.TestRun 失败（统计翻倍）。
