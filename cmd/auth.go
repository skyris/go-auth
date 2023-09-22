package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/skyris/auth-server/internal/database"
	"github.com/skyris/auth-server/internal/database/postgres"
	httpLayer "github.com/skyris/auth-server/internal/http"
	useCaseUser "github.com/skyris/auth-server/internal/usecase/user"
)

func main() {
	conn, err := postgres.New(postgres.NewSettings())
	if err != nil {
		log.Panicln(err)
	}

	defer conn.Pool.Close()

	storage, err := database.New(conn.Pool, database.Options{})
	if err != nil {
		log.Panicln(err)
	}
	ucUser := useCaseUser.New(storage, useCaseUser.Options{})

	httpListener := httpLayer.New(ucUser, httpLayer.Options{})

	go func() {
		if err := httpListener.Run(); err != nil {
			log.Panicln(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
