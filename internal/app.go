package app

import (
	"log"
	"os"
)

type app struct {
	name string
	log  log.Logger
}

func Run() {
	App := &app{
		"size-calc",
		*log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	App.log.Printf("hello world!")
}
