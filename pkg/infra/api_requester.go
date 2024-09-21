package infra

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ApiRequester struct {
	Url string
}

func NewApiRequester(url string) *ApiRequester {
	return &ApiRequester{Url: url}
}

func (a *ApiRequester) MakeRequest() (interface{}, error) {
	var apiResponse interface{}
	cl := http.Client{}

	req, err := http.NewRequest("GET", a.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("not found result for the zip code")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return apiResponse, nil
}
