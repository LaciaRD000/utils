package http

import (
	http "github.com/bogdanfinn/fhttp"
	tlsClient "github.com/bogdanfinn/tls-client"
)

func Post(url string, headers http.Header, json map[string]string, options []tlsClient.HttpClientOption) (*http.Response, error) {
	client, err := tlsClient.NewHttpClient(tlsClient.NewLogger(), options...)
	if err != nil {
		return nil, err
	}
	payload := setPayload(json)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, url, payload)
	req.Header = headers
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
