package main

import (
	server2 "github.com/AkbarFikri/Mooistudio-Technical-Test/internal/config/server"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	app := gin.New()
	server := server2.New(app)
	server.Run()
}
