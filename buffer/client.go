package buffer

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// URL is the base URL of the Buffer.com API
	URL = "https://api.bufferapp.com/1"
)

// Client is the Buffer.com API client
type Client struct {
	AccessToken string
	URL         string
}

// Update represents a Buffer.com update
type Update struct {
	ID             string
	Text           string
	ProfileService string `json:"profile_service"`
	TextFormatted  string `json:"text_formatted"`
}

// Updates represents a set of Buffer.com updates
type Updates []Update

// NewClient creates a new Buffer.com API client
func NewClient(accessToken string) *Client {
	return &Client{URL: URL, AccessToken: accessToken}
}

// Push creates a new status update for one or more profiles
func (c *Client) Push(text string, profileIDs []string) (Updates, error) {
	params := url.Values{}
	params.Set("text", text)
	for _, p := range profileIDs {
		params.Add("profile_ids[]", p)
	}

	res, err := c.sendPOST("updates/create", params)
	if err != nil {
		return nil, err
	}

	var response struct {
		Success          bool
		BufferCount      int
		BufferPercentage int
		Updates          Updates
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, err
	}

	if response.Success == false {
		return nil, errors.New("Unable to create a new update (Buffer is likely full)")
	}

	return response.Updates, nil
}

func (c *Client) sendPOST(resource string, params url.Values) ([]byte, error) {
	req, err := http.PostForm(c.URL+"/"+resource+".json?access_token="+c.AccessToken, params)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return ioutil.ReadAll(req.Body)
}
