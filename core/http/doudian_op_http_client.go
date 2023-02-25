package http

import (
	"bytes"
	"crypto/tls"
	"doudian.com/open/sdk_golang/errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var clientMap sync.Map

type DoudianOpHttpClient struct {
	httpClient *http.Client
}

type DoudianOpHttpRequest struct {
	Url     string
	Params  map[string]string
	Headers map[string]string
	Body    string
}

type DoudianOpHttpResponse struct {
	Body string
}

func newHttpClient(timeout int64) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives:   false,
		MaxIdleConns:        50000,
		MaxIdleConnsPerHost: 50000,
		IdleConnTimeout:     60 * time.Second,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}

	return &http.Client{Transport: transport, Timeout: time.Duration(timeout) * time.Millisecond}
}

func (client *DoudianOpHttpClient) Post(httpRequest *DoudianOpHttpRequest) (*DoudianOpHttpResponse, error) {
	u, err := url.Parse(httpRequest.Url)
	if err != nil {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.HttpError, err.Error())
	}
	if len(httpRequest.Params) > 0 {
		query := u.Query()
		for k, v := range httpRequest.Params {
			query.Add(k, v)
		}
		u.RawQuery = query.Encode()
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBufferString(httpRequest.Body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "Content-Type: application/json")
	if len(httpRequest.Headers) > 0 {
		for k, v := range httpRequest.Headers {
			req.Header.Set(k, v)
		}
	}
	httpResp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.HttpError, err.Error())
	}

	if httpResp.StatusCode != http.StatusOK {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.HttpError, fmt.Sprintf("http code = %d", httpResp.StatusCode))
	}
	bs, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.HttpError, err.Error())
	}

	return &DoudianOpHttpResponse{Body: string(bs)}, nil
}

func GetHttpClient(timeout int64) *DoudianOpHttpClient {
	doudianHttpClient, ok := clientMap.Load(timeout)

	//并发高的情况下，可能导致httpClient初始化两次
	if !ok || doudianHttpClient == nil {
		httpClient := newHttpClient(timeout)
		newDoudianOpHttpClient := &DoudianOpHttpClient{httpClient: httpClient}
		clientMap.Store(timeout, newDoudianOpHttpClient)
		return newDoudianOpHttpClient
	}
	return doudianHttpClient.(*DoudianOpHttpClient)
}
