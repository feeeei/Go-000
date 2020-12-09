## 作业
基于 `errgroup` 实现一个 `http server` 的启动和关闭 ，以及 `linux signal` 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 个人理解
感觉用 `errgroup` 不是很契合这个场景，因为 `errgroup` 的 `Wait()` 需要等待所有任务执行完毕后才会返回，而作业场景更适合的是一个退出，剩下的也跟着退出。 最合适的方案感觉应该是使用 `channel` 的扇入扇出去实现，类似下述这样
```go
pipegroup.Go(func(in chan error, out chan error) {
	// 如果我挂了，我往 out 发 error，告知别人不用跑了我挂了
	// 如果别人挂了，应该往 in 发 error，告知我不用跑了
})
```
思路就是实现一个 `pipegroup`，其中维护一个叫 control 的统一 channel 用来接收 job 异常时传递过来的 error。然后为每个 job 维护一个 channel，当 control 接收到 error 时，使用扇出的方式通知每一个 job。

## 代码实现
- 使用 errgroup 实现的版本：[errgroup版本](https://github.com/feeeei/Go-000/blob/main/Week03/errgroup.md)

- 使用自定义 pipegroup 实现的版本：[pipegroup版本](https://github.com/feeeei/Go-000/blob/main/Week03/pipegroup.md)