package server

import (
	"github.com/labstack/echo/v4"
	_ "github.com/paulojunior/code-challange/cmd/api/docs"
	contractIntegration "github.com/paulojunior/code-challange/contract/integration"
	contract "github.com/paulojunior/code-challange/contract/service"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SetRoutes(e *echo.Echo, musicService contract.MusicService, spotifyIntegration contractIntegration.SpotifyIntegration) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	g := e.Group("/api/v1")

	g.Use(JWTMiddleware)

	g.GET("/isrc/:isrc", func(c echo.Context) error {
		return HandlerGetByISRC(c, musicService)
	})

	g.GET("/artist/:name", func(c echo.Context) error {
		return HandleGetByArtistName(c, musicService)
	})

	g.POST("/insert-track", func(c echo.Context) error {
		return HandlerInsertTrack(c, musicService, spotifyIntegration)
	})
}
