package main

import (
	"context"
	"logserver/internal/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	a, err := app.New()
	if err != nil {
		panic(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	complete := make(chan struct{})

	go func() {
		<-sig
		cancel()
		if err := a.Stop(complete); err != nil {
			panic(err)
		}
	}()

	if err := a.Run(ctx); err != nil {
		panic(err)
	}
}
