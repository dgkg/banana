package concert

import (
	"net/http"
	"net/url"
)

type SDKConcert struct {
	key string
	cli *http.Client
}

type Concert struct {
	Mbid           string `json:"mbid"`
	Name           string `json:"name"`
	SortName       string `json:"sortName"`
	Disambiguation string `json:"disambiguation"`
	URL            string `json:"url"`
}

func (sdk *SDKConcert) GetConcertByID(uuid string) (*Concert, error) {
	// create the full url
	url := BaseURL + "Concert/" + uuid
	// try to exec the request
	var a Concert
	err := execGet(sdk.cli, url, sdk.key, &a)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func (sdk *SDKConcert) SearchConcerts(query map[string]string) ([]Concert, error) {
	if query == nil {
		query = make(map[string]string)
	}
	// create the full url
	urltoCall := BaseURL + "search/Concerts"
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
	var as []Concert
	err = execGet(sdk.cli, u.String(), sdk.key, &as)
	if err != nil {
		return nil, err
	}
	return as, nil
}
