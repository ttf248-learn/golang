## 笔记

相当于开发语言学习的 hello world，文档地址：[https://go-zero.dev/docs/concepts/overview](https://go-zero.dev/docs/concepts/overview)

* handler: 路由注册的地方
* logic: 业务逻辑处理的地方
* svc: 传递数据库等外部对象配置

运行以后发现的东西，除开请求日志，项目默认会输出运行的资源信息，包含硬件资源占用、节点的 qps 统计信息

{"@timestamp":"2024-04-24T16:29:48.666+08:00","caller":"stat/usage.go:61","content":"CPU: 0m, MEMORY: Alloc=0.9Mi, TotalAlloc=0.9Mi, Sys=8.1Mi, NumGC=0","level":"stat"}
{"@timestamp":"2024-04-24T16:29:48.666+08:00","caller":"load/sheddingstat.go:61","content":"(api) shedding_stat [1m], cpu: 0, total: 0, pass: 0, drop: 0","level":"stat"}
{"@timestamp":"2024-04-24T16:30:08.011+08:00","caller":"stat/metrics.go:210","content":"(ch13_go_zero_greet-api) - qps: 0.0/s, drops: 0, avg time: 0.0ms, med: 0.0ms, 90th: 0.0ms, 99th: 0.0ms, 99.9th: 0.0ms","level":"stat"}
