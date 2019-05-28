package main

import (
	"os"
	"os/signal"
	"syscall"
	"translator/internal/dao"
	"translator/internal/views"
	"translator/pkg/translator"
)

func main() {
	memoryDAO := dao.NewMemoryDAO()

	t := translator.New(memoryDAO)

	cliMemory := views.NewCLIMemory(t)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)
	go func() {
		<-stop
		cliMemory.Outro()
		os.Exit(0)
	}()

	cliMemory.Intro()

	for {
		cliMemory.CallToAction()
		cliMemory.ResolveAction()
	}
}
