package main

import (
	"crudserver/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a, err := app.New()
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	complete := make(chan struct{})

	go func() {
		<-sig
		if err := a.Stop(complete); err != nil {
			panic(err)
		}
	}()

	if err := a.Run(); err != nil {
		panic(err)
	}
}
