// +build cli

package internal

import (
	"io"
	"os"
	"os/signal"
	"syscall"
)

type CLI interface {
	Intro()
	Outro()
	CallToAction()
	ResolveAction() (err error)
}

func RunCLI(cliMemory CLI) {
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
		err := cliMemory.ResolveAction()
		if err == io.EOF {
			return
		}
	}
}
