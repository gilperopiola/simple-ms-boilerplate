package main

import (
	"flag"
	"log"
	"os"

	"github.com/gilperopiola/simple-ms-boilerplate/config"
)

var cfg config.MyConfig
var rtr MyRouter

func main() {
	env := flag.String("env", "local", "local / dev / prod")
	flag.Parse()

	cfg.Setup(*env)
	rtr.Setup()

	log.Println("server started")
	rtr.Run(":" + os.Getenv("PORT"))
}
