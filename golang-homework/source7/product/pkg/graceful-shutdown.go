package web

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
)

type GracefulShutdownAction interface {
	ShutdownFilter(f Filter) Filter
	WaitingForSVR(ctx context.Context, server... *Server) error
	GraceShutdown(server... *Server)
}

type ShutDownAction struct {
	// 访问handler的数量
	couter int64
	// 通知主动关闭的信号量
	closing int64
	// 完成关闭信号量
	final chan int

}

func NewShutdown() *ShutDownAction {
	return &ShutDownAction{
		couter: 0,
		closing: 0,
		final: make(chan int, 1),
	}
}

func (s *ShutDownAction) ShutdownBuilder(f Filter) Filter{
	return func(c *Context) {
		closing := atomic.LoadInt64(&s.closing)
		// 拒绝新请求
		if closing > 0 {
			c.W.WriteHeader(http.StatusServiceUnavailable)
		}
		atomic.AddInt64(&s.couter, 1)
		f(c)
		atomic.AddInt64(&s.couter, -1)
		// 等待请求结束
		counter := atomic.LoadInt64(&s.couter)
		fmt.Println(counter)
		if closing > 0 && counter == 0{
			s.final <- 1
		}
	}
}

func (s *ShutDownAction) WaitingForSVR(ctx context.Context, server... *Server) error{
	atomic.AddInt64(&s.closing, 1)
	closing := atomic.LoadInt64(&s.closing)
	counter := atomic.LoadInt64(&s.couter)
	if closing > 0 && counter == 0{
		s.final <- 1
	}
	select {
	case <-s.final:
		for _, svr := range server {
			svr.Shutdown()
		}
		return nil
	case <-ctx.Done():
		return errors.New("Timeout")
	}
}

func (s *ShutDownAction)GraceShutdown(server... *Server) {
	sysSiginal := make(chan os.Signal, 1)

	signal.Notify(sysSiginal, syscall.SIGTERM, syscall.SIGINT)
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	select {
	case <-sysSiginal:
		if err := s.WaitingForSVR(ctx, server...); err != nil {
			fmt.Println(err)
		}
	}

}

