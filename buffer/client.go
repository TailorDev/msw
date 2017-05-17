package buffer

import (
	"encoding/json"
	"errors"
	"fmt"
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
	Day            string
	ProfileService string `json:"profile_service"`
	TextFormatted  string `json:"text_formatted"`
}

// Updates represents a set of Buffer.com updates
type Updates []Update

// Profile represents a Buffer.com profile
type Profile struct {
	ID                string
	Service           string
	FormattedUsername string `json:"formatted_username"`
}

// NewClient creates a new Buffer.com API client
func NewClient(accessToken string) *Client {
	return &Client{URL: URL, AccessToken: accessToken}
}

// Push creates a new status update for one or more profiles
func (c *Client) Push(text string, profileIDs []string) (Updates, error) {
	var u Updates
	params := url.Values{}
	params.Set("text", text)
	for _, p := range profileIDs {
		params.Add("profile_ids[]", p)
	}

	res, err := c.post("updates/create", params)
	if err != nil {
		return u, err
	}

	var response struct {
		Success          bool
		BufferCount      int
		BufferPercentage int
		Updates          Updates
		Message          string
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return u, err
	}

	if response.Success == false {
		return u, errors.New(response.Message)
	}

	u = response.Updates

	return u, nil
}

func (c *Client) GetProfile(profileID string) (Profile, error) {
	var p Profile
	res, err := c.get(fmt.Sprintf("profiles/%s", profileID))
	if err != nil {
		return p, err
	}

	err = json.Unmarshal(res, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (c *Client) GetPendingUpdates(profileID string) (Updates, error) {
	var u Updates
	res, err := c.get(fmt.Sprintf("profiles/%s/updates/pending", profileID))
	if err != nil {
		return u, err
	}

	var response struct {
		Total   int
		Updates Updates
	}

	err = json.Unmarshal(res, &response)
	if err != nil {
		return u, err
	}

	u = response.Updates

	return u, nil
}

func (c *Client) get(res string) ([]byte, error) {
	req, err := http.Get(fmt.Sprintf("%s/%s.json?access_token=%s", c.URL, res, c.AccessToken))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return ioutil.ReadAll(req.Body)
}

func (c *Client) post(res string, params url.Values) ([]byte, error) {
	req, err := http.PostForm(fmt.Sprintf("%s/%s.json?access_token=%s", c.URL, res, c.AccessToken), params)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	return ioutil.ReadAll(req.Body)
}
