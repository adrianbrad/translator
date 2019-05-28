package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"translator"
	"translator/internal/dao"
	"translator/internal/views"
)

func main() {
	port := flag.Int("p", 8080, "Sets the servers port")
	flag.Parse()
	memoryDAO := dao.NewMemoryDAO()

	t := translator.New(memoryDAO)

	webView := views.NewWebView(t)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: webView,
	}

	go func() {
		err := server.ListenAndServe()
		fmt.Println(err.Error())
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)
	<-stop
	ctx := context.TODO()
	server.Shutdown(ctx)
}
