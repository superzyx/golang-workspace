package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	// io.WriteString(w, "ok")
	query := r.URL.Query()
	id := query.Get("name")
	fmt.Println("----", id)

	// io.WriteString(w, id)
	w.Header().Set("Name", id)
	w.Header().Set("Version", runtime.Version())
	w.WriteHeader(200)
	io.WriteString(w, "hello worlds\n")
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Header.Get(`X-Real-Ip`))
	fmt.Println(r.Header.Get(`X-Forwarded-For`))

	io.WriteString(w, "状态码: 200\n")
	io.WriteString(w, "client访问地址："+formatIP(r))

}

func formatIP(r *http.Request) string {
	remoteIp := r.RemoteAddr
	if ip := r.Header.Get(`X-Real-Ip`); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get(`X-Forwarded-For`); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}

	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}
	return remoteIp
}
