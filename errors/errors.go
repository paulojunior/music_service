package errors

import (
	"fmt"
)

func NewErrHTTPReqCreation(err error) error {
	return fmt.Errorf("error creating http request: %s", err.Error())
}

func NewErrAccessToken(err error) error {
	return fmt.Errorf("error getting Spotify access token: %s", err.Error())
}

func NewErrAPICall(err error) error {
	return fmt.Errorf("error making call to Spotify API: %s", err.Error())
}

func NewErrReadingResponseBody(err error) error {
	return fmt.Errorf("error reading response body: %s", err.Error())
}

func NewErrJSONParsing(err error) error {
	return fmt.Errorf("error parsing JSON: %s", err.Error())
}
