## 作业
基于 `errgroup` 实现一个 `http server` 的启动和关闭 ，以及 `linux signal` 信号的注册和处理，要保证能够 一个退出，全部注销退出。

## 代码实现
```go
func main() {
	group, cancelCtx := errgroup.WithContext(context.Background())
	group.Go(func() error {
		return serveHTTP(cancelCtx)
	})
	group.Go(func() (err error) {
		return serveSignal(cancelCtx)
	})

	// 两个都挂了之后，运行收尾 shutdown()
	if err := group.Wait(); err != nil {
		shutdown()
	}
}

// serveHTTP 启动业务
func serveHTTP(ctx context.Context) error {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	// 模拟一个HTTP挂了的情况
	http.HandleFunc("/boom", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "boom!!!")
		server.Shutdown(context.Background())
	})
	go func() {
		<-ctx.Done()
		server.Shutdown(context.Background())
	}()
	return server.ListenAndServe()
}

// serveSignal 监听关闭信号
func serveSignal(ctx context.Context) error {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c:
		return fmt.Errorf("Receive shutdown signal")
	case <-ctx.Done():
		log.Println("兄弟死了，那我也不活了")
		return nil
	}
}

// shutdown 执行一些收尾工作，最多等1秒
func shutdown() {
	finish := make(chan struct{})
	go func() {
		log.Println("Starting shutdown...")
		// close_conn
		// flush_log
		// clean_cache
		// ......
		time.Sleep(time.Millisecond * 200) // 模拟后台处理收尾任务200ms
		log.Println("Finished shutdown")
		finish <- struct{}{}
	}()
	select {
	case <-finish:
		return
	case <-time.After(time.Second * 1):
		return
	}
}
```