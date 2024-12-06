package concert

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type SDKAPI struct {
	Artists *SDKArtist
	Concert *SDKConcert
}

func New(key string) *SDKAPI {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}

	return &SDKAPI{
		Artists: &SDKArtist{
			key: key,
			cli: client,
		},
	}
}

const BaseURL = "https://api.setlist.fm/rest/1.0/"

func execGet(cli *http.Client, url, key string, payload interface{}) error {
	// init the request GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// add the header mandatory elements
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", key)
	req.Header.Add("X-Idempotency-Key", uuid.NewString())

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	// never forget to close the body response
	defer res.Body.Close()
	// read the body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// check the status code of the response
	if res.StatusCode != http.StatusOK {
		var errResp ErrConcertResponse
		// bind the response into the struct givent in param
		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return err
		}
		return errResp
	}
	// bind the response into the struct givent in param
	err = json.Unmarshal(body, payload)
	if err != nil {
		return err
	}
	return nil
}
