package fugle

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host     string
	version  string
	apiToken string
}

type ClientOption struct {
	host     string
	version  string
	apiToken string
}

func NewFugleClient(opts ClientOption) (Client, error) {
	client := &Client{
		host:     opts.host,
		version:  opts.version,
		apiToken: opts.apiToken,
	}
	if client.host == "" {
		client.host = "api.fugle.tw"
	}
	if client.version == "" {
		client.version = "v0.2"
	}
	if client.apiToken == "" {
		return *client, errors.New("please provide fugle api token")
	}

	return *client, nil
}

func (cli *Client) callAPI(endpoint string, params url.Values) (*Response, error) {
	targetURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	params.Add("apiToken", cli.apiToken)
	targetURL.RawQuery = params.Encode()

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "*/*")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	response := &Response{}
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (cli *Client) Chart(symbolID string, oddLot bool) (*Response, error) {
	baseUrl := fmt.Sprintf("https://%s/realtime/%s/intraday/chart", cli.host, cli.version)
	params := url.Values{}
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	return cli.callAPI(baseUrl, params)
}

func (cli *Client) Quote(symbolID string, oddLot bool) (*Response, error) {
	baseUrl := fmt.Sprintf("https://%s/realtime/%s/intraday/quote", cli.host, cli.version)
	params := url.Values{}
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	return cli.callAPI(baseUrl, params)
}

func (cli *Client) Meta(symbolID string, oddLot bool) (*Response, error) {
	baseUrl := fmt.Sprintf("https://%s/realtime/%s/intraday/meta", cli.host, cli.version)
	params := url.Values{}
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	return cli.callAPI(baseUrl, params)
}

func (cli *Client) Dealts(symbolID string, oddLot bool, limit int64, offset int64) (*Response, error) {
	baseUrl := fmt.Sprintf("https://%s/realtime/%s/intraday/dealts", cli.host, cli.version)
	params := url.Values{}
	params.Add("symbolId", symbolID)
	params.Add("oddLot", strconv.FormatBool(oddLot))
	params.Add("limit", strconv.FormatInt(limit, 10))
	params.Add("offset", strconv.FormatInt(offset, 10))
	return cli.callAPI(baseUrl, params)
}
