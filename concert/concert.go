package concert

import (
	"encoding/json"
	"io"
	"net/http"
)

type SDKAPI struct {
	key string
}

func New(key string) *SDKAPI {
	return &SDKAPI{key: key}
}

const BaseURL = "https://api.setlist.fm/rest/1.0/"

func (sdk *SDKAPI) execGet(url string, payload interface{}) error {
	// init the request GET
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	// add the header mandatory elements
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", sdk.key)
	// exec the request
	res, err := http.DefaultClient.Do(req)
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
	// bind the response into the struct givent in param
	err = json.Unmarshal(body, payload)
	if err != nil {
		return err
	}
	return nil
}
