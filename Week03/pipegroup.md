## 使用自定义 pipegroup 实现的版本

```go
# pipegroup
type Group struct {
	workers []chan error
	control chan error
}

func New() *Group {
	return &Group{
		workers: make([]chan error, 0),
		control: make(chan error),
	}
}

func (g *Group) Wait() error {
	err := <-g.control
	// 使用扇出的方式，通知每一个 worker 可以停了
	for _, worker := range g.workers {
		worker <- err
	}
	return err
}

func (g *Group) Go(f func(in chan error, out chan error)) {
	worker := make(chan error, 1)
	g.workers = append(g.workers, worker)
	go f(worker, g.control)
}
```

```go
func main() {
	group := pipegroup.New()
	group.Go(serveHTTP)
	group.Go(serveSignal)

	if err := group.Wait(); err != nil {
		log.Println("pipe_group error:", err)
		shutdown()
	}
}

func serveHTTP(in chan error, out chan error) {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	go func() {
		<-in
		server.Shutdown(context.Background())
	}()
	log.Println("Starting http service...")
	if err := server.ListenAndServe(); err != nil {
		out <- fmt.Errorf("Http service stop")
	}
}

func serveSignal(in chan error, out chan error) {
	c := make(chan os.Signal, 1)
	defer close(c)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-c:
		out <- fmt.Errorf("Receive shutdown signal")
	case <-in:
		return
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
		log.Println("End of timeout")
		return
	}
}
```