## 作业
基于 `errgroup` 实现一个 `http server` 的启动和关闭 ，以及 `linux signal` 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 个人理解
感觉用 `errgroup` 不是很契合这个场景，因为 `errgroup` 的 `Wait()` 需要等待所有任务执行完毕后才会返回，而作业场景更适合的是一个退出，剩下的也跟着退出。 最合适的方案感觉应该是使用 `channel` 的扇入扇出去实现，类似下述这样
```go
errgroup.Go(func(receive chan error, output chan error) {
	// 如果我挂了，我往 output 发 error，告知别人不用跑了我挂了
	// 如果别人挂了，应该往 receive 发 error，告知我有其他人挂了
})
```

## 代码实现
> //TODO 老师这次作业留的时间太紧张了，先实现完成作业的版本，扇入扇出版本等等补
```go
func main() {
	stop, cancel := begin()
	group, _ := errgroup.WithContext(stop)
	group.Go(func() error {
		defer cancel()
		return serveHTTP(stop)
	})
	group.Go(func() (err error) {
		defer cancel()
		return serveSignal(stop)
	})

	// 两个都挂了之后，运行收尾shutdown()
	if err := group.Wait(); err != nil {
		shutdown()
	}
}

// 名字叫 begin 不太合适，一时半会儿想不到更好的...
// 把 context 的 cancel 用 once 包一层，保证 cancel 只执行一次
func begin() (context.Context, func()) {
	stop, cancel := context.WithCancel(context.Background())
	var once sync.Once
	c := func() {
		once.Do(cancel)
	}
	return stop, c
}

// serveHTTP 启动业务
func serveHTTP(ctx context.Context) (err error) {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	// 模拟一个HTTP挂了的情况
	http.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "boom!!!")
		server.Shutdown(context.Background())
	})
	go func() { // 异步监听关闭信号
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()
	err = server.ListenAndServe()
	log.Printf("http server stop: %v", err)
	return
}

// serveSignal 监听关闭信号
func serveSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c:
		log.Println("收到关闭信号")
		return fmt.Errorf("Receive shutdown signal")
	case <-ctx.Done():
		log.Println("兄弟死了，那我也不活了")
		return nil
	}
}

// shutdown 执行一些收尾工作，最多等1秒
func shutdown() {
	down, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	finish := make(chan struct{})
	go func() {
		// close_conn
		// flush_log
		// clean_cache
		// ......
		time.Sleep(time.Millisecond * 200) // 模拟后台处理收尾任务200ms
		finish<-struct{}{}
	}()
	select {
	case <-finish:
		return
	case <-down.Done():
		return
	}
}
```