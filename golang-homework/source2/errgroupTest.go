package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func welcome(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("welcome!!"))
}

func main() {
	ctx := context.Background()
	ctx, can := context.WithCancel(ctx)
	g, ctx := errgroup.WithContext(ctx)

	sig := make(chan os.Signal, 1)
	ch := make(chan int, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	svr := &http.Server{Addr: ":8989"}

	g.Go(func() error{
		http.HandleFunc("/welcome", welcome)
		if err := svr.ListenAndServe(); err != nil {
			select {
			case ch <- 1:
				fmt.Println("server 异常")
			default:
			}
			return err
		}
		return nil
	})

	g.Go(func() error {
		<- ctx.Done()
		err := svr.Shutdown(ctx)
		fmt.Println("级联删除server")
		return errors.New(fmt.Sprintf("主动关闭 %d", err))

	})

	g.Go(func() error {
		select {
		case <- sig:
		case <- ctx.Done():
			fmt.Println("级联删除")
		}

		select {
		case ch <- 1:
		default:
		}
		return errors.New("server shutdown")
	})

	select {
	case <-ch :
		ch <- 1
		can()
	}
	if err := g.Wait();err != nil {
		fmt.Println(err)
	}
}
