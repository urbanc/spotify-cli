package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"spotify/pkg/model"
)

type APIInterface interface {
	Play() error
	Pause() error
}

type API struct {
	token string
}

func NewAPI(token string) *API {
	return &API{token}
}

func (s *API) Play() error {
	return s.call("PUT", "/me/player/play")
}

func (s *API) Pause() error {
	return s.call("PUT", "/me/player/pause")
}

func (s *API) call(method string, endpoint string) error {
	url := "https://api.spotify.com/v1" + endpoint

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token))

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Success
	if res.StatusCode == http.StatusNoContent {
		return nil
	}

	// Error
	spotifyErr := new(model.SpotifyError)
	if err := json.NewDecoder(res.Body).Decode(spotifyErr); err != nil {
		return err
	}

	return errors.New(spotifyErr.Error.Message)
}
