package app

import "github.com/nunoonu/mail-consumer/internal/handlers"

func ProvideApp(httpServer *handlers.HTTPService, kafkaServer *handlers.KafkaService) *App {
	return &App{
		HTTPServer:  *httpServer,
		KafkaServer: *kafkaServer,
	}
}

type App struct {
	HTTPServer  handlers.HTTPService
	KafkaServer handlers.KafkaService
}
