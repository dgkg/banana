package concert

import (
	"net/http"
	"net/url"
)

type SDKArtist struct {
	key string
	cli *http.Client
}

type Artist struct {
	Mbid           string `json:"mbid"`
	Name           string `json:"name"`
	SortName       string `json:"sortName"`
	Disambiguation string `json:"disambiguation"`
	URL            string `json:"url"`
}

type ArtistListResponse struct {
	Type         string   `json:"type"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Page         int      `json:"page"`
	Total        int      `json:"total"`
	Artist       []Artist `json:"artist"`
}

func (sdk *SDKArtist) GetByID(uuid string) (*Artist, error) {
	// create the full url
	url := BaseURL + "artist/" + uuid
	// try to exec the request
	var a Artist
	err := execGet(sdk.cli, url, sdk.key, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (sdk *SDKArtist) Search(query map[string]string) ([]Artist, error) {
	if query == nil {
		query = make(map[string]string)
	}
	// create the full url
	urltoCall := BaseURL + "search/artists"
	u, err := url.Parse(urltoCall)
	if err != nil {
		return nil, err
	}
	q := url.Values{}
	for k, v := range query {
		// query[k] = url.QueryEscape(v)
		q.Add(k, v)
	}
	if len(q) > 0 {
		u.RawQuery = q.Encode()
	}
	// try to exec the request
	var response ArtistListResponse
	err = execGet(sdk.cli, u.String(), sdk.key, &response)
	if err != nil {
		return nil, err
	}

	return response.Artist, nil
}
