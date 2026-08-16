# BUG_REPRO

## Bug 是什么
包厢容量判断把“已订满”误判为可订；预约存储去锁后 check-then-act 出现数据竞争；批量预约的成功计数无锁导致丢失更新；预约统计把包厢数也重复加进去。

## 如何触发
`go test -race ./...`

## 错误信息
- model.TestCanBook 失败（容量已满仍可订）。
- service.TestBatchReserve 失败（超卖，reserved=5 期望 4），`-race` 报 DATA RACE。
- store.TestReserveRejectsWhenFull 失败（满员仍预约成功）。
- worker.TestRun 失败（总数翻倍）。
