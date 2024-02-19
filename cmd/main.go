package main

import (
	"github.com/nunoonu/mail-consumer/cmd/app"
	"log/slog"
	"sync"
)

func main() {
	ap := app.InitializeApp()
	slog.Info("App is initialized")
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func(a app.App) {
		defer wg.Done()
		go a.KafkaServer.Start()
		a.HTTPServer.Start(a.KafkaServer.Close)
	}(*ap)

	wg.Wait()
}
