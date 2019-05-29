// +build web

package internal

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func RunWeb(web http.Handler) {
	port := flag.Int("p", 8080, "Sets the servers port")
	flag.Parse()

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: web,
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
	os.Exit(0)
}
