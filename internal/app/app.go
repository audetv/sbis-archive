package app

import (
	"log"
	"os"

	"github.com/audetv/sbis-archive/internal/config"
	"github.com/audetv/sbis-archive/internal/entities/document"
	"github.com/audetv/sbis-archive/internal/storage/manticore"
)

type App struct {
	manticoreStorage *document.Entries
}

func NewApp(cfg *config.Config) *App {

	var manticoreStorage document.Entries
	manticoreStorage = *NewEntriesStorage(cfg.ManticoreIndex.Name)

	return &App{
		manticoreStorage: &manticoreStorage,
	}
}

func NewEntriesStorage(index string) *document.Entries {
	var storage document.StorageInterface

	manticoreClient, err := manticore.New(index)
	if err != nil {
		log.Printf("failed to initialize manticore client for index %v, %v", index, err)
		os.Exit(1)
	}

	storage = manticoreClient

	return document.NewAnswerStorage(storage)
}
