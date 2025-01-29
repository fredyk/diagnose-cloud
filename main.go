package main

import (
	"log"

	"github.com/fredyk/diagnose-cloud/server"
)

func main() {

	log.Fatal(server.InitServer())

}
