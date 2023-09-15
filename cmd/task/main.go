package main

import (
	"github.com/sirupsen/logrus"

	"gitlab.com/zapirus/task/config"
	"gitlab.com/zapirus/task/internal/pkg/app"
)

func main() {
	conf, err := config.NewConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	server := app.New(conf)

	server.Run()
}
