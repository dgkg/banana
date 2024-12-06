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

func (sdk *SDKArtist) GetArtistByID(uuid string) (*Artist, error) {
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

func (sdk *SDKArtist) SearchArtists(query map[string]string) ([]Artist, error) {
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
	var as []Artist
	err = execGet(sdk.cli, u.String(), sdk.key, &as)
	if err != nil {
		return nil, err
	}
	return as, nil
}
