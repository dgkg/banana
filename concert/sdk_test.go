package concert

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bytedance/sonic"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fastjson"
)

func TestExecCassyGet_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"type":"artistList","itemsPerPage":10,"page":1,"total":1,"artist":[{"mbid":"123","name":"Test Artist","sortName":"Artist, Test","disambiguation":"","url":"http://example.com"}]}`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	payload := &ArtistListResponse{}
	err := execStdWithDecoderGet(client, server.URL, "test-key", payload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if payload.Type != "artistList" {
		t.Errorf("expected type 'artistList', got %s", payload.Type)
	}
	if len(payload.Artist) != 1 {
		t.Errorf("expected 1 artist, got %d", len(payload.Artist))
	}
	if payload.Artist[0].Name != "Test Artist" {
		t.Errorf("expected artist name 'Test Artist', got %s", payload.Artist[0].Name)
	}
}

func TestExecCassyGet_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":400,"status":"error","message":"Bad Request","timestamp":"2023-10-10T10:00:00Z"}`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	payload := &ArtistListResponse{}
	err := execStdWithDecoderGet(client, server.URL, "test-key", payload)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if errResp, ok := err.(ErrConcertResponse); ok {
		if errResp.Code != 400 {
			t.Errorf("expected error code 400, got %d", errResp.Code)
		}
		if errResp.Message != "Bad Request" {
			t.Errorf("expected error message 'Bad Request', got %s", errResp.Message)
		}
	} else {
		t.Fatalf("expected ErrConcertResponse, got %T", err)
	}
}

func TestExecWithFastjsonGet_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`[{"mbid":"d76b8372-a70d-4c48-a26e-06d625114620","name":"A. Billie Free","sortName":"A. Billie Free","disambiguation":"","url":"https://www.setlist.fm/setlists/a-billie-free-23f690a3.html"},{"mbid":"28bb2507-de99-4197-be13-ad1af6977b08","name":"Billie Anderson","sortName":"Anderson, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-anderson-3bdf5058.html"},{"mbid":"de61176b-ef33-4931-9b7e-a21f129babd5","name":"Billie Anthony","sortName":"Anthony, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-anthony-1bd799c8.html"},{"mbid":"eb25edde-8df8-49f7-8c4e-21b5c6ffacf7","name":"Aquasky vs. Masterblaster feat. Billie Godfrey","sortName":"Aquasky vs. Masterblaster feat. Godfrey, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/aquasky-vs-masterblaster-feat-billie-godfrey-5bdafbbc.html"},{"mbid":"5d06fe54-485a-4a07-b506-5f6f719448cb","name":"Billie Joe Armstrong","sortName":"Armstrong, Billie Joe","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-joe-armstrong-bd689b2.html"},{"mbid":"868b71bf-b814-47bf-8288-6d6680c5a6d3","name":"Billie Joe Armstrong \u0026 Penelope Houston","sortName":"Armstrong, Billie Joe \u0026 Houston, Penelope","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-joe-armstrong-and-penelope-houston-2bd6447a.html"},{"mbid":"d083325a-d7ae-4f0e-b01f-642a75bfcdb7","name":"Billie Joe + Norah","sortName":"Armstrong, Billie Joe + Jones, Norah","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-joe-norah-13da7135.html"},{"mbid":"5cdb5f93-fc52-4490-aabe-cf8498c6197d","name":"Louis Armstrong \u0026 Billie Holiday","sortName":"Armstrong, Louis \u0026 Holiday, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-and-billie-holiday-4bdf37c2.html"},{"mbid":"9f5131dc-544b-472f-ae38-4894c72c5611","name":"Louis Armstrong \u0026 Billie Holiday with Sy Oliver’s Orchestra","sortName":"Armstrong, Louis \u0026 Holiday, Billie with Oliver, Sy and His Orchestra","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-and-billie-holiday-with-sy-olivers-orchestra-1bcbe504.html"},{"mbid":"7db1b6bb-769b-4be3-9129-ed8b5d5eb769","name":"Louis Armstrong and Billie Holiday","sortName":"Armstrong, Louis and Holiday, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-and-billie-holiday-5bde6b1c.html"},{"mbid":"4a4e878e-20dd-49b0-8c49-6325be70f40e","name":"Louis Armstrong and Billie Holiday with Sy Oliver and His Orchestra","sortName":"Armstrong, Louis and Holiday, Billie with Oliver, Sy and His Orchestra","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-and-billie-holiday-with-sy-oliver-and-his-orchestra-bfaedde.html"},{"mbid":"5ea47db6-8ac1-49a0-887d-3ca08ea6b7f0","name":"Louis Armstrong with Billie Holiday","sortName":"Armstrong, Louis with Holiday, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-with-billie-holiday-3bdc3898.html"},{"mbid":"b6155db8-cade-462e-8515-b97717fd182f","name":"Louis Armstrong, Billie Holiday \u0026 Sy Oliver and His Orchestra","sortName":"Armstrong, Louis, Holiday, Billie \u0026 Oliver, Sy and His Orchestra","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-billie-holiday-and-sy-oliver-and-his-orchestra-33f04861.html"},{"mbid":"76b94112-ae55-4415-a514-f4b755cdc2c1","name":"Louis Armstrong, Billie Holiday \u0026 Sy Oliver’s Orchestra","sortName":"Armstrong, Louis, Holiday, Billie \u0026 Oliver, Sy and His Orchestra","disambiguation":"","url":"https://www.setlist.fm/setlists/louis-armstrong-billie-holiday-and-sy-olivers-orchestra-43f0435b.html"},{"mbid":"c87a5429-2237-4abe-81b4-760da939215e","name":"Joshua Atchley, Billie Ray Fingers \u0026 Bruce Fingers","sortName":"Atchley, Joshua, Fingers, Billie Ray \u0026 Fingers, Bruce","disambiguation":"","url":"https://www.setlist.fm/setlists/joshua-atchley-billie-ray-fingers-and-bruce-fingers-2bf18092.html"},{"mbid":"fadda754-2bca-49fc-9e60-dc37656deb7c","name":"Joshua Atchley, Bruce Fingers \u0026 Billie Ray Fingers","sortName":"Atchley, Joshua, Fingers, Bruce \u0026 Fingers, Billie Ray","disambiguation":"","url":"https://www.setlist.fm/setlists/joshua-atchley-bruce-fingers-and-billie-ray-fingers-13f58579.html"},{"mbid":"37961821-8a4f-40c0-b87e-571b8788e6bd","name":"Bess Atwell, Billie Marten \u0026 Ellie Mason","sortName":"Atwell, Bess, Marten, Billie \u0026 Voka Gentle","disambiguation":"","url":"https://www.setlist.fm/setlists/bess-atwell-billie-marten-and-ellie-mason-43fa5b47.html"},{"mbid":"7d65a3eb-7394-4422-861b-0d211d7163f1","name":"Billie Barnum","sortName":"Barnum, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/billie-barnum-23d5fc37.html"},{"mbid":"6db32b1b-d4a4-47fc-8645-bd924cad1430","name":"Sara Barone, Bruce Fingers \u0026 Billie Ray Fingers","sortName":"Barone, Sara, Fingers, Bruce \u0026 Fingers, Billie Ray","disambiguation":"","url":"https://www.setlist.fm/setlists/sara-barone-bruce-fingers-and-billie-ray-fingers-6bf58ab6.html"},{"mbid":"b16c94fc-a7b8-40f8-bd9e-749737942a0f","name":"Tony Bennett duet with Billie Holiday","sortName":"Bennett, Tony duet with Holiday, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/tony-bennett-duet-with-billie-holiday-1bf3ed90.html"},{"mbid":"079670cd-595b-4906-89d3-59caccc21751","name":"Tony Bennett feat Billie Holiday","sortName":"Bennett, Tony feat Holiday, Billie","disambiguation":"","url":"https://www.setlist.fm/setlists/tony-bennett-feat-billie-holiday-73c75ac5.html"},{"mbid":"a2fd6cac-8a48-4152-bf0c-d5068b9c8690","name":"Bigger Than Billie","sortName":"Bigger Than Billie","disambiguation":"Canadian","url":"https://www.setlist.fm/setlists/bigger-than-billie-43c43be3.html"},{"mbid":"f31be735-020d-489c-9103-ccbb12b4754c","name":"BILLIE","sortName":"BILLIE","disambiguation":"appears on Basile Palace's track","url":"https://www.setlist.fm/setlists/billie-73e24a29.html"},{"mbid":"71f31cfe-2882-48ed-ab53-fb7c3273f264","name":"Billie","sortName":"Billie","disambiguation":"French pop girl duet","url":"https://www.setlist.fm/setlists/billie-43d577e7.html"},{"mbid":"add55b02-e95d-4b49-810d-576d37999cf4","name":"Billie","sortName":"Billie","disambiguation":"saxophonist","url":"https://www.setlist.fm/setlists/billie-2bc140e6.html"},{"mbid":"ea22b3b7-8666-4ae1-ab71-66565b756b3e","name":"Billie","sortName":"Billie","disambiguation":"Brooklyn, NY singer Billie Brown","url":"https://www.setlist.fm/setlists/billie-6bdadec6.html"},{"mbid":"ca88941a-24d2-4c45-8e42-7d07f373da90","name":"Billie","sortName":"Billie","disambiguation":"Techno","url":"https://www.setlist.fm/setlists/billie-4bfa6b6e.html"},{"mbid":"4cea60ef-a211-424d-89d0-c115303dfac9","name":"Billie","sortName":"Billie","disambiguation":"Billie Wrede","url":"https://www.setlist.fm/setlists/billie-2bcfecc2.html"},{"mbid":"0fe3bf80-e6a6-4748-a529-e5fbfc62e0b0","name":"Billie","sortName":"Billie","disambiguation":"French singer Amélie Lacaf","url":"https://www.setlist.fm/setlists/billie-33c7e439.html"},{"mbid":"cd2c68df-748a-440a-8fa2-350d904785a6","name":"Billie","sortName":"Billie","disambiguation":"UK singer/actor Billie Piper","url":"https://www.setlist.fm/setlists/billie-73d7f241.html"}]`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	payload := &ArtistListResponse{}
	err := execFastjsonGet(client, server.URL, "test-key", payload)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if payload.Type != "artistList" {
		t.Errorf("expected type 'artistList', got %s", payload.Type)
	}
	if len(payload.Artist) != 1 {
		t.Errorf("expected 1 artist, got %d", len(payload.Artist))
	}
	if payload.Artist[0].Name != "Test Artist" {
		t.Errorf("expected artist name 'Test Artist', got %s", payload.Artist[0].Name)
	}
}

func TestExecWithFastjsonGet_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"code":400,"status":"error","message":"Bad Request","timestamp":"2023-10-10T10:00:00Z"}`))
	}))
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	payload := &ArtistListResponse{}
	err := execFastjsonGet(client, server.URL, "test-key", payload)
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	if errResp, ok := err.(ErrConcertResponse); ok {
		if errResp.Code != 400 {
			t.Errorf("expected error code 400, got %d", errResp.Code)
		}
		if errResp.Message != "Bad Request" {
			t.Errorf("expected error message 'Bad Request', got %s", errResp.Message)
		}
	} else {
		t.Fatalf("expected ErrConcertResponse, got %T", err)
	}
}

func BenchmarkExecCassyGet(b *testing.B) {
	server := mockServerForBench()
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	for i := 0; i < b.N; i++ {
		payload := &ArtistListResponse{}
		err := execStdWithDecoderGet(client, server.URL, "test-key", payload)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkExecFastjsonGet(b *testing.B) {
	server := mockServerForBench()
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	for i := 0; i < b.N; i++ {
		payload := &ArtistListResponse{}
		err := execFastjsonGet(client, server.URL, "test-key", payload)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkExecJsoniterGet(b *testing.B) {
	server := mockServerForBench()
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	for i := 0; i < b.N; i++ {
		payload := &ArtistListResponse{}
		err := execJsoniterGet(client, server.URL, "test-key", payload)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkExecSonicGet(b *testing.B) {
	server := mockServerForBench()
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	for i := 0; i < b.N; i++ {
		payload := &ArtistListResponse{}
		err := execSonicGet(client, server.URL, "test-key", payload)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkExecStdGet(b *testing.B) {
	server := mockServerForBench()
	defer server.Close()

	client := &http.Client{Timeout: Timeout}
	for i := 0; i < b.N; i++ {
		payload := &ArtistListResponse{}
		err := execStdGet(client, server.URL, "test-key", payload)
		if err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func mockServerForBench() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(mockArtistPayload())
	}))
}

func mockArtistPayload() []byte {
	return []byte(`{"type":"artistList","itemsPerPage":10,"page":1,"total":1,"artist":[{"mbid":"123","name":"Test Artist","sortName":"Artist, Test","disambiguation":"","url":"http://example.com"}]}`)
}

func BenchmarkParseStd(b *testing.B) {
	data := mockArtistPayload() // Exemple de données JSON
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		if err := json.Unmarshal(data, &result); err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkParseJsoniter(b *testing.B) {
	data := mockArtistPayload() // Exemple de données JSON
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		if err := jsoniter.Unmarshal(data, &result); err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkParseFastjson(b *testing.B) {
	data := mockArtistPayload() // Exemple de données JSON
	for i := 0; i < b.N; i++ {
		var p fastjson.Parser
		if _, err := p.ParseBytes(data); err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkParseGoJson(b *testing.B) {
	data := mockArtistPayload() // Exemple de données JSON
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		if err := gojson.Unmarshal(data, &result); err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}

func BenchmarkParseSonic(b *testing.B) {
	data := mockArtistPayload() // Exemple de données JSON
	for i := 0; i < b.N; i++ {
		var result map[string]interface{}
		if err := sonic.Unmarshal(data, &result); err != nil {
			b.Fatalf("expected no error, got %v", err)
		}
	}
}
