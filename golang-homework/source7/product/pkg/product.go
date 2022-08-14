package web

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var(
	data = make(chan int, 10)
)

func healthz(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "200")

}


func working(writer http.ResponseWriter, request *http.Request) {

	for i:=0 ; i <10 ; i++ {
		time.Sleep(1*time.Second)
		log.Println("INFO: process: ", i)
	}
	log.Println("INFO: end in one process")
	//data <- 1
}

func gracefullyExit(ctx context.Context, timeOut time.Duration) {
	select {
	case <- ctx.Done():
		log.Println("exiting %s ...", ctx.Err())
	case <- time.After(timeOut):
		log.Println("exiting...")
	}

}


func main() {
	ch := make(chan os.Signal, 1)


	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/worker", working)
	server := &http.Server{
		Addr: ":8888",
		Handler: mux,
		//WriteTimeout: 5*time.Second,
	}


	go func(){
		if err := server.ListenAndServe(); err!= nil && err != http.ErrServerClosed{
			log.Fatalln("Fatal: ", err)
		}
	}()

	log.Println("INFO: start the server")




	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ch :
		log.Println("graceful exit...")
		ctx ,cen := context.WithTimeout(context.Background(), 5*time.Second)
		defer cen()
		gracefullyExit(ctx, 2*time.Second)
	}

	log.Println("INFO: stop the server")
}

