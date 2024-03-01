package main

import (
	"log"

	"github.com/stdyum/api-studyplaces/internal"
)

func main() {
	log.Fatalf("error launching web server %s", internal.App().Error())
}
