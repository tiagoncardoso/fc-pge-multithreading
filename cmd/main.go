package main

import (
	"fmt"
	"github.com/tiagoncardoso/fc/pge/multithreading/config"
	"github.com/tiagoncardoso/fc/pge/multithreading/pkg/infra"
	"log/slog"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Print("\033[H\033[2J")

	ch1 := make(chan interface{})
	ch2 := make(chan interface{})

	conf, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		panic("Missing zip code")
	}

	url1 := buildApiUrl(conf.ApiRequest1URL, os.Args[1])
	url2 := buildApiUrl(conf.ApiRequest2URL, os.Args[1])

	go makeRequest(url1, ch1, 0)
	go makeRequest(url2, ch2, 1)

	select {
	case resp1 := <-ch1:
		slog.Info("Faster Response from", "Api URL", url1, "resp:", resp1)
	case resp2 := <-ch2:
		slog.Info("Faster Response from", "Api URL", url2, "resp:", resp2)
	case <-time.After(time.Second * time.Duration(conf.RequestTimeout)):
		slog.Error("Timeout making request")
	}

	// TODO: format json response: Slog
	// TODO: create documentation
}

func makeRequest(url string, ch chan interface{}, delay int) {
	time.Sleep(time.Duration(delay) * time.Second)

	requester := infra.NewApiRequester(url)
	resp, err := requester.MakeRequest()
	if err != nil {
		slog.Error("Error making request", "msg", err)
		<-ch
	}

	ch <- resp
}

func buildApiUrl(originalUrl string, zipCode string) string {
	return strings.Replace(originalUrl, "<<zip>>", zipCode, -1)
}
