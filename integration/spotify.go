package integration

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/paulojunior/code-challange/errors"
	"github.com/spf13/viper"
)

type TracksSpotifyDTO struct {
	Tracks struct {
		Items []Item  `json:"items"`
		Next  *string `json:"next"`
	} `json:"tracks"`
}

type Item struct {
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Artists    []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Album struct {
		Images []struct {
			Url    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"images"`
	} `json:"album"`
}

type SpotifyIntegration struct{}

func NewSpotifyIntegration() *SpotifyIntegration {
	return &SpotifyIntegration{}
}

func (i SpotifyIntegration) GetTrackByISRC(ISRC string) (Item, error) {
	client := http.Client{}
	var allTracks []Item

	offset := 0
	limit := 20

	for {
		url := fmt.Sprintf("%s/search?q=isrc:%s&type=track&offset=%d&limit=%d", viper.GetString("spotify.api_url"), ISRC, offset, limit)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Item{}, errors.NewErrHTTPReqCreation(err)
		}

		acessToken, err := i.getSpotifyToken()
		if err != nil {
			return Item{}, err
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", acessToken))

		var resp *http.Response
		for attempt := 1; attempt <= 3; attempt++ {
			resp, err = client.Do(req)
			if err == nil {
				break
			}
		}

		if resp == nil {
			return Item{}, errors.NewErrAPICall(err)
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return Item{}, errors.NewErrReadingResponseBody(err)
		}

		var result TracksSpotifyDTO
		err = json.Unmarshal(body, &result)
		if err != nil {
			return Item{}, errors.NewErrJSONParsing(err)
		}

		allTracks = append(allTracks, i.GetMostPopularTrack(result.Tracks.Items))

		if result.Tracks.Next == nil {
			break
		}

		offset += limit
	}

	return i.GetMostPopularTrack(allTracks), nil
}

func (i SpotifyIntegration) GetMostPopularTrack(items []Item) Item {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Popularity > items[j].Popularity
	})

	if len(items) > 0 {
		return items[0]
	}

	return Item{}
}

func (i SpotifyIntegration) getSpotifyToken() (string, error) {
	clientID, clientSecret := viper.GetString("spotify.client_id"), viper.GetString("spotify.client_secret")
	authURL := viper.GetString("spotify.auth_url")

	authHeader := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Basic "+authHeader)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", errors.NewErrJSONParsing(err)
	}

	accessToken, ok := response["access_token"].(string)
	if !ok {
		return "", errors.NewErrJSONParsing(err)
	}

	return accessToken, nil
}
