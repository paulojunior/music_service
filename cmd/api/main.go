package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/paulojunior/code-challange/database"
	"github.com/paulojunior/code-challange/integration"
	"github.com/paulojunior/code-challange/repository"
	"github.com/paulojunior/code-challange/server"
	"github.com/paulojunior/code-challange/service"
	"github.com/spf13/viper"
)

// @title Swagger Music Track API
// @version 1.0
// @description This is a music track API.

// @contact.name Paulo Ferreira
// @contact.email jr@live.at

// @host localhost:3000
// @BasePath /
func main() {
	log.Println("Starting api server...")

	viper.SetConfigName("config")
	viper.AddConfigPath("../../config/")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	db := database.NewPostgresDatabase()

	e := echo.New()
	server.SetRoutes(e, service.NewMusicService(repository.NewMusicRepository(db)), integration.NewSpotifyIntegration())
	e.Logger.Fatal(e.Start(":8080"))
}
