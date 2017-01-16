package twitter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"net/http"
	"net/url"
)

// BaseURL is the base URL for making twitter API calls
var BaseURL = "https://api.twitter.com/1.1"

// Client sends requests that handle JSON encoding and decoding for request and response bodies
type Client struct {
	HTTPclient   *http.Client
	BaseURL      string
	apiKey       string
	AccessToken  string
	clientSecret string
	Logger       *logrus.Logger
}

// New returns a new *HTTPClient
func New(baseURL, apiKey, secret string, logger *logrus.Logger) *Client {
	return &Client{
		HTTPclient:   &http.Client{},
		BaseURL:      baseURL,
		apiKey:       apiKey,
		clientSecret: secret,
		Logger:       logger,
	}
}

// Feed get a user's twitter feed
func Feed() {
	print("tweet")
}

// Access gets the access token
func (c *Client) Access() {
	creds := fmt.Sprintf("%s:%s", c.apiKey, c.clientSecret)
	encodedCreds := base64.URLEncoding.EncodeToString([]byte(creds))

	body := url.Values{}
	body.Add("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", bytes.NewBufferString(body.Encode()))
	if err != nil {
		fmt.Printf(err.Error())
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", encodedCreds))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8.")

	resp, err := c.HTTPclient.Do(req)
	if err != nil {
		fmt.Printf(err.Error())
	}

	defer resp.Body.Close()
	auth := AuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(&auth)
	if err != nil {
		fmt.Printf(err.Error())
	}
	c.AccessToken = auth.AccessToken

	fmt.Printf("twitter access dest%#+v\n", auth)
}

// Call invokes the http request
func (c *Client) Call(req *http.Request, dest interface{}) (*http.Response, error) {
	// Add a header to request the response in json
	req.Header.Add("response_content_type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := c.HTTPclient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(dest)

	if err != nil {
		fmt.Printf("err 1%#+v\n", err.Error())
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest { // StatusBadRequest=400, everything above is not OK
		c.Logger.Errorf("HTTP client: Endpoint: %s, Method: %s, StatusCode: %d", req.URL, req.Method, resp.StatusCode)
	}

	return resp, err
}
