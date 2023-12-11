package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	contractIntegration "github.com/paulojunior/code-challange/contract/integration"
	contract "github.com/paulojunior/code-challange/contract/service"
	"github.com/paulojunior/code-challange/entity"
)

// @Summary Get music by ISRC
// @Description Retrieve music information by its ISRC (International Standard Recording Code)
// @Tags music
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC of the music" format(isrc)
// @Success 200 {array} object "Successfully retrieved music information"
// @Failure 400 {string} string "ISRC not found"
// @Failure 404 {string} string "Music not found"
// @Failure 500 {string} string "Internal server error"
// @Router /music/{isrc} [get]
func HandlerGetByISRC(c echo.Context, musicService contract.MusicService) error {
	isrc := c.Param("isrc")

	if isrc == "" {
		return c.String(http.StatusBadRequest, "ISRC not found")
	}

	musics, err := musicService.FindByISRC(isrc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error to fetch music")
	}

	return c.JSON(http.StatusOK, musics)
}

// @Summary Get music by artist name
// @Description Retrieve music information by the name of the artist
// @Tags music
// @Accept json
// @Produce json
// @Param name path string true "Name of the artist" example("John Doe")
// @Success 200 {array} object "Successfully retrieved music information"
// @Failure 400 {string} string "Name cannot be empty"
// @Failure 404 {string} string "Music not found"
// @Failure 500 {string} string "Internal server error"
// @Router /music/artist/{name} [get]
func HandleGetByArtistName(c echo.Context, musicService contract.MusicService) error {
	name := c.Param("name")

	if name == "" {
		return c.String(http.StatusBadRequest, "name cannot be empty")
	}

	musics, err := musicService.FindByArtistName(name)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error to fetch music")
	}

	return c.JSON(http.StatusOK, musics)
}

// @Summary Insert music track metadata
// @Description Retrieve track information from Spotify based on the provided ISRC
// @Tags tracks
// @Accept json
// @Produce json
// @Param isrc path string true "ISRC of the track"
// @Success 201 {object} string "Track inserted successfully"
// @Failure 400 {string} string "ISRC not found"
// @Failure 404 {string} string "Track not found"
// @Failure 500 {string} string "Internal server error"
// @Router /tracks/{isrc} [post]
func HandlerInsertTrack(c echo.Context, musicService contract.MusicService, spotifyIntegration contractIntegration.SpotifyIntegration) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return c.String(http.StatusBadRequest, "Failed to parse request body")
	}

	isrc, ok := requestBody["isrc"].(string)
	if !ok || isrc == "" {
		return c.String(http.StatusBadRequest, "ISRC not found in the JSON body")
	}

	musics, err := musicService.FindByISRC(isrc)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Error to fetch track")
	}

	if len(musics) > 0 {
		return c.JSON(http.StatusConflict, "The song already exists in the database.")
	}

	result, err := spotifyIntegration.GetTrackByISRC(isrc)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	var artists []string
	for _, artist := range result.Artists {
		artists = append(artists, artist.Name)
	}

	var largestImageURL string
	largestWidth := 0

	for _, image := range result.Album.Images {
		if image.Width > largestWidth {
			largestWidth = image.Width
			largestImageURL = image.Url
		}
	}

	music := entity.Music{
		ISRC:    isrc,
		Title:   result.Name,
		Artists: artists,
	}

	if largestImageURL != "" {
		music.ImageURL = largestImageURL
	}

	err = musicService.Insert(music)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, music)
}
