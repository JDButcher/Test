package app

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	db "testProject/internal/database/postgresql"
	operation "testProject/internal/services"
	"testProject/pkg/http"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})

	postgresqlConfig := db.Config{}
	err := envconfig.Process("", &postgresqlConfig)
	if err != nil {
		log.Fatal(err)
	}

	serverConfig := http.ServerConfig{}
	err = envconfig.Process("", &serverConfig)
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewDBRepository(context.Background(), &postgresqlConfig)

	operationService := operation.NewOperationService(repository)

	log.Infof(
		http.StartServer(&serverConfig, operationService),
	)

}
