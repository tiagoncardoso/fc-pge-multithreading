package main

import (
	"github.com/tiagoncardoso/fc/pge/multithreading/config"
	"github.com/tiagoncardoso/fc/pge/multithreading/pkg/infra"
	"log/slog"
	"time"
)

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	conf, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	go makeRequest(conf.ApiRequest1URL, ch1, 1)
	go makeRequest(conf.ApiRequest2URL, ch2, 0)

	select {
	case resp1 := <-ch1:
		slog.Info("Faster Response from", "Api URL", conf.ApiRequest1URL, "resp:", resp1)
	case resp2 := <-ch2:
		slog.Info("Faster Response from", "Api URL", conf.ApiRequest2URL, "resp:", resp2)
	case <-time.After(time.Second * time.Duration(conf.RequestTimeout)):
		slog.Error("Timeout making request")
	}
}

func makeRequest(url string, ch chan<- interface{}, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)

	requester := infra.NewApiRequester(url)
	resp, err := requester.MakeRequest()
	if err != nil {
		slog.Error("Error making request", "msg", err)
	}

	ch <- resp
}
