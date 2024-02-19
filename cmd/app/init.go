package app

import (
	"github.com/nunoonu/mail-consumer/helpers"
	"github.com/nunoonu/mail-consumer/internal/core/usecases"
	"github.com/nunoonu/mail-consumer/internal/handlers"
	"github.com/nunoonu/mail-consumer/internal/repositories"
	"log/slog"
	"os"
)

func InitializeApp() *App {
	setLogLevel()
	params := helpers.NewMailKafkaParams()
	kr := helpers.NewKafka(params)

	mailParams := repositories.NewMailParams()
	mailRepo := repositories.NewMailRepository(mailParams)
	fileUsc := usecases.NewFileUseCase(mailRepo)
	fileHdl := handlers.NewFileHandler(fileUsc)
	kafkaServ := handlers.NewKafkaService(kr, fileHdl)

	router := handlers.NewRouter()
	httpServParams := handlers.NewHTTPServiceParams()
	httpServ := handlers.NewHTTPService(httpServParams, router)
	return ProvideApp(httpServ, kafkaServ)
}

func setLogLevel() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(l)
}
