package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key string
	value string
}

var theTests = []struct {
	name string
	url string
	method string
	params []postData
	expextedStatusCode int
} {
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"mr", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search-availailability post", "/search-availability", "POST", []postData{
		{key: "start", value: "2023-08-02"},
		{key: "end", value: "2023-08-04"},
	}, http.StatusOK},
	{"search-availability-json post", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2023-08-02"},
		{key: "end", value: "2023-08-04"},
	}, http.StatusOK},
	{"make-reservation post", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "555-55-5555"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expextedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expextedStatusCode, resp.StatusCode )
			}
		} else {
			values := url.Values{}
			for _, x := range e.params{
				values.Add(x.key, x.value)
			}
			resp, err := ts.Client().PostForm(ts.URL + e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expextedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expextedStatusCode, resp.StatusCode )
			}
		}
	}
}