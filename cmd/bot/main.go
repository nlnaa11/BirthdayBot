package main

import (
	"log"
	"os"

	"github.com/nlnaa11/BirthdayBot/internal/clients"
	"github.com/nlnaa11/BirthdayBot/internal/config"
	"github.com/nlnaa11/BirthdayBot/internal/model/messages"
)

func main() {
	files, err := os.ReadDir("data/")
	if err != nil {
		log.Fatal("Failed to read a data folder: ", err)
	}

	var cnf *config.Service

	if len(files) > 1 {
		cnf, err = config.RecoverConfig()
	}
	else {
		cnf, err = config.InitConfig()
	}

	if err != nil {
		log.Fatal("config init failed: ", err)
	}

	tgClient, err := clients.NewTgClient(cnf)
	if err != nil {
		log.Fatal("tg client init failed", err)
	}

	helper := messages.NewModelHelper(config)

	msgModel := messages.NewModel(tgClient, helper)

	tgClient.ListenUpdates(msgModel)
}
