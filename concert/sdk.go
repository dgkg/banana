package concert

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fastjson"
)

type SDKAPI struct {
	Artists *SDKArtist
	Concert *SDKConcert
}

const Timeout = 10 * time.Second

func New(key string) *SDKAPI {
	tr := &http.Transport{
		MaxIdleConns:          1000,
		MaxIdleConnsPerHost:   100,
		MaxConnsPerHost:       500,
		IdleConnTimeout:       Timeout,
		TLSHandshakeTimeout:   Timeout,
		ExpectContinueTimeout: 1 * time.Second,
	}
	client := &http.Client{Transport: tr}

	return &SDKAPI{
		Artists: &SDKArtist{
			key: key,
			cli: client,
		},
	}
}

const (
	BaseURL      = "https://api.setlist.fm/rest/1.0/"
	MaxBodyBytes = 1 * 1024 * 1024 // 1 MB
)

var isWithFastjson = true

func execGet(cli *http.Client, url, key string, payload any) error {
	if isWithFastjson {
		return execFastjsonGet(cli, url, key, payload)
	}
	return execStdWithDecoderGet(cli, url, key, payload)
}

func execStdWithDecoderGet(cli *http.Client, url, key string, payload any) error {
	// create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	// create the request with the context timeout.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// add the header mandatory elements.
	addRequestHeader(req, key)

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	// never forget to close the body response.
	defer res.Body.Close()

	// check the content type response.
	err = validateJSONResponseHeader(res)
	if err != nil {
		return err
	}

	// limit the size of the response body
	limitedReader := &io.LimitedReader{R: res.Body, N: MaxBodyBytes}
	decoder := json.NewDecoder(limitedReader)

	// check the status code of the response.
	// if it's not 200, then it should be an error.
	if res.StatusCode != http.StatusOK {
		var errResp ErrConcertResponse
		return decodeBodyResponse(limitedReader, decoder, true, &errResp)
	}

	// decode the response body.
	return decodeBodyResponse(limitedReader, decoder, false, payload)
}

func execStdGet(cli *http.Client, url, key string, payload any) error {
	// create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	// create the request with the context timeout.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// add the header mandatory elements.
	addRequestHeader(req, key)

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	// never forget to close the body response.
	defer res.Body.Close()

	// check the content type response.
	err = validateJSONResponseHeader(res)
	if err != nil {
		return err
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// limit the size of the response body
	if len(data) > MaxBodyBytes {
		return fmt.Errorf("response body too large")
	}

	// check the status code of the response.
	// if it's not 200, then it should be an error.
	if res.StatusCode != http.StatusOK {
		var errResp ErrConcertResponse
		return json.Unmarshal(data, &errResp)
	}

	// decode the response body.
	return json.Unmarshal(data, payload)
}

func execFastjsonGet(cli *http.Client, url, key string, payload any) error {
	// create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	// create the request with the context timeout.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// add the header mandatory elements.
	addRequestHeader(req, key)

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check the content type response.
	err = validateJSONResponseHeader(res)
	if err != nil {
		return err
	}

	// limit the size of the response body
	limitedReader := &io.LimitedReader{R: res.Body, N: MaxBodyBytes}
	body, err := io.ReadAll(limitedReader)
	if err != nil {
		return err
	}

	// check if the response body was truncated
	if limitedReader.N == 0 {
		return fmt.Errorf("response body too large")
	}

	// parse the JSON response using fastjson
	var p fastjson.Parser
	v, err := p.ParseBytes(body)
	if err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	// handle the parsed JSON (example for a specific structure)
	// you need to adapt this part to your specific payload structure
	err = parsePayload(v, payload)
	if err != nil {
		return err
	}

	return nil
}

func parsePayload(v *fastjson.Value, payload interface{}) error {
	switch payloadConcreat := payload.(type) {
	case *ArtistListResponse:
		payloadConcreat.Type = string(v.GetStringBytes("type"))
		payloadConcreat.ItemsPerPage = v.GetInt("itemsPerPage")
		payloadConcreat.Page = v.GetInt("page")
		payloadConcreat.Total = v.GetInt("total")
		artists := v.GetArray("artist")
		payloadConcreat.Artist = make([]Artist, len(artists))
		for i, artist := range artists {
			if err := parsePayload(artist, &payloadConcreat.Artist[i]); err != nil {
				return err
			}
		}
		return nil

	case *Artist:
		*payloadConcreat = Artist{
			Mbid:           string(v.GetStringBytes("mbid")),
			Name:           string(v.GetStringBytes("name")),
			SortName:       string(v.GetStringBytes("sortName")),
			Disambiguation: string(v.GetStringBytes("disambiguation")),
			URL:            string(v.GetStringBytes("url")),
		}
		return nil

	case *Concert:
		*payloadConcreat = Concert{
			Mbid:           string(v.GetStringBytes("mbid")),
			Name:           string(v.GetStringBytes("name")),
			SortName:       string(v.GetStringBytes("sortName")),
			Disambiguation: string(v.GetStringBytes("disambiguation")),
			URL:            string(v.GetStringBytes("url")),
		}
		return nil

	case *[]Concert:
		concerts := v.GetArray("concert")
		*payloadConcreat = make([]Concert, len(concerts))
		for i, concert := range concerts {
			if err := parsePayload(concert, &(*payloadConcreat)[i]); err != nil {
				return err
			}
		}
		return nil

	case *ErrConcertResponse:
		payloadConcreat.Code = v.GetInt("code")
		payloadConcreat.Status = string(v.GetStringBytes("status"))
		payloadConcreat.Message = string(v.GetStringBytes("message"))
		payloadConcreat.Timestamp = string(v.GetStringBytes("timestamp"))
		return nil
	default:
		// implement the parsing for other payload types.
		return fmt.Errorf("unsupported payload type given: %T", payload)
	}
}

func addRequestHeader(req *http.Request, key string) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", key)
	req.Header.Add("X-Idempotency-Key", uuid.NewString())
}

func validateJSONResponseHeader(res *http.Response) error {
	if contentType := res.Header.Get("Content-Type"); contentType != "application/json" {
		return fmt.Errorf("unexpected content type: %s", contentType)
	}

	return nil
}

func execJsoniterGet(cli *http.Client, url, key string, payload any) error {

	// create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	// create the request with the context timeout.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// add the header mandatory elements.
	addRequestHeader(req, key)

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	// never forget to close the body response.
	defer res.Body.Close()

	// check the content type response.
	err = validateJSONResponseHeader(res)
	if err != nil {
		return err
	}

	// limit the size of the response body
	limitedReader := &io.LimitedReader{R: res.Body, N: MaxBodyBytes}
	decoder := jsoniter.NewDecoder(limitedReader)

	// check the status code of the response.
	// if it's not 200, then it should be an error.
	if res.StatusCode != http.StatusOK {
		var errResp ErrConcertResponse
		return decodeBodyResponse(limitedReader, decoder, true, &errResp)
	}

	// decode the response body.
	return decodeBodyResponse(limitedReader, decoder, false, payload)
}

func execSonicGet(cli *http.Client, url, key string, payload any) error {

	// create a context with a timeout.
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()

	// create the request with the context timeout.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	// add the header mandatory elements.
	addRequestHeader(req, key)

	// exec the request
	res, err := cli.Do(req)
	if err != nil {
		return err
	}

	// never forget to close the body response.
	defer res.Body.Close()

	// check the content type response.
	err = validateJSONResponseHeader(res)
	if err != nil {
		return err
	}

	// read the response body.
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// check the status code of the response.
	// if it's not 200, then it should be an error.
	if res.StatusCode != http.StatusOK {
		var errResp ErrConcertResponse
		err = sonic.Unmarshal(data, &errResp)
		if err != nil {
			return err
		}
		return errResp
	}

	// decode the response body.
	return sonic.Unmarshal(data, payload)
}

func decodeBodyResponse(limit *io.LimitedReader, decode Decoder, withError bool, payload any) error {
	// decode the response body.
	err := decode.Decode(payload)
	if err != nil {
		return err
	}

	// check if the response body was truncated.
	if limit.N == 0 {
		return fmt.Errorf("response body too large")
	}

	// check if the response should be an error.
	if withError {
		return payload.(ErrConcertResponse)
	}

	return nil
}

type Decoder interface {
	Decode(obj any) error
}
