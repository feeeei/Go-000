## 作业
基于 `errgroup` 实现一个 `http server` 的启动和关闭 ，以及 `linux signal` 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 个人理解
感觉用 `errgroup` 不是很契合这个场景，因为 `errgroup` 的 `Wait()` 需要等待所有任务执行完毕后才会返回，而作业场景更适合的是一个退出，剩下的也跟着退出。 

所以我认为，最契合的 `errgroup` 应该是一种使用 `channel` 扇入扇出模式实现的，当异步的 Job 发生异常时，向 `errgroup` 发 error（扇入操作），`errgroup` 接收到后通知每一个Job可以停了（扇出操作）。
```go
type PipeGroup struct {
	control chan error   // Job 有异常了发给我
	workers []chan error // 异常广播给每一个 Worker
}

pipegroup.Go(func(in chan error, out chan error) {
	// 如果我挂了，我往 out 发 error，告知别人不用跑了我挂了
	// 如果别人挂了，应该往 in 发 error，告知我不用跑了
})
```

## 代码实现
- 使用 errgroup，context 实现的版本：[errgroup，context版本](https://github.com/feeeei/Go-000/blob/main/Week03/errgroup.md)

- 使用自定义 pipegroup 实现的版本：[pipegroup版本](https://github.com/feeeei/Go-000/blob/main/Week03/pipegroup.md)