package config

import (
	"log"
	"myforum/internal/models"
)

type Application struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	ForumModel *models.ForumModel
}

var App = &Application{}
