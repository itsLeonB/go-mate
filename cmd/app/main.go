package main

import (
	"github.com/itsLeonB/go-mate/internal/delivery/http/server"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	srv := server.SetupServer()
	srv.Serve()
}
