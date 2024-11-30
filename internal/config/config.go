package config

import (
	"log"
	modules "myforum/internal/modules"
)

type Application struct {
	ErrorLog   *log.Logger
	InfoLog    *log.Logger
	ForumModel *modules.ForumModel
}

var App = &Application{}
