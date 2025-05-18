package main

import (
	"log"
	"os"
)

type app struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (a *app) handleError(message string, err error) {
    if err != nil {
        a.errorLogger.Printf("%s: %v", message, err)
        os.Exit(1)
    }
}