package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	web "web-go/pkg"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type SignUpReq struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Password string `json:"password"`
}

type CommonResponse struct {
	Code int `json:"status_code"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func registry(ctx *web.Context) {
	writer := ctx.W
	request := ctx.R
	req := &SignUpReq{}
	res := &CommonResponse{}

	info, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Errorf("ERROR: %s", err)
		return
	}
	err = json.Unmarshal(info, req)
	if err != nil {
		fmt.Errorf("ERROR: format error: %s", err)
	}
	writer.WriteHeader(http.StatusOK)

	response, err := json.Marshal(req)
	res.Code = http.StatusOK
	res.Msg = "jjjsds"
	res.Data = string(response)
	resByte, _ := json.Marshal(&res)
	_, _ = writer.Write(resByte)
}

func healthz(c *web.Context) {
	c.W.WriteHeader(http.StatusOK)
}

func promhttpHandle(c *web.Context) {
	promethusSVC := promhttp.Handler()
	promethusSVC.ServeHTTP(c.W, c.R)
}

func main() {
	sd := web.NewShutdown()
	web.Register()
	s := web.NewServer("newOne", sd.ShutdownBuilder, web.NewMetrics())
	s.Route("GET","/registry", registry)
	s.Route("GET", "/healthz", healthz)
	s.Route("GET", "/metrics", promhttpHandle)
	go s.Start(":8888")
	sd.GraceShutdown(s, )
}




