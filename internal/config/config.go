package config


import "log"

type Application struct {
	ErrorLog *log.Logger
	InfoLog *log.Logger
}

var App = &Application{}
