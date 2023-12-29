package main

import (
	"errors"
	"io"
	"os"
	"strconv"

	"net/http"
	"net/http/httputil"
	"net/url"

	"encoding/json"
)

var (
	targetHost = os.Getenv("TARGET_HOST")
)

func main() {
	proxy, err := newProxy()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", proxyRequestHandler(proxy))
	http.ListenAndServe(":8000", nil)
}

func newProxy() (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	return httputil.NewSingleHostReverseProxy(url), nil
}

func proxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		rewrittenRequest, err := rewriteRequest(request)
		if err != nil {
			http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
			return
		}

		proxy.ServeHTTP(responseWriter, rewrittenRequest)
	}
}

func rewriteRequest(request *http.Request) (*http.Request, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, errors.New("error reading request body")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.New("error parsing request body")
	}

	rewrittenRequest, err := http.NewRequest(request.Method, targetHost, nil)
	if err != nil {
		return nil, errors.New("error creating proxy request")
	}

	rewrittenRequest.URL.Path = "/index.php/apps/maps/api/1.0/devices"
	rewrittenRequest.URL.RawQuery = buildQuery(request, data).Encode()

	rewrittenRequest.Header.Add("Authorization", request.Header.Get("Authorization"))

	return rewrittenRequest, nil
}

func buildQuery(request *http.Request, body map[string]interface{}) url.Values {
	query := make(url.Values)
	query.Add("user_agent", request.Header.Get("X-Limit-D"))
	query.Add("lat", getAsFloat(body["lat"]))
	query.Add("lng", getAsFloat(body["lon"]))
	query.Add("timestamp", getAsInt(body["tst"]))
	return query
}

func getAsFloat(data interface{}) string {
	value := data.(float64)
	return strconv.FormatFloat(value, 'f', 10, 64)
}

func getAsInt(data interface{}) string {
	value := int(data.(float64))
	return strconv.Itoa(value)
}
