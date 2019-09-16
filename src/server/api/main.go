package main

import (
	"log"
	"os"

	"github.com/k-ueki/app2/src/server"
)

func main() {
	log.SetFlags(log.Ldate + log.Ltime + log.Lshortfile)
	log.SetOutput(os.Stdout)

	s := server.NewServer()
	s.Init()
}
