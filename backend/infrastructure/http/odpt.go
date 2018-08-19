package http

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/yuki-toida/refodpt/backend/config"
)

type Client struct {
	BaseURL     string
	ConsumerKey string
}

func NewOdptClient() *Client {
	return &Client{
		BaseURL:     config.Config.OdptURL,
		ConsumerKey: config.Config.ConsumerKey,
	}
}

func (c *Client) Get(path string, target interface{}) error {
	return c.GetBy(path, target, map[string]string{})
}

func (c *Client) GetBy(path string, target interface{}, args map[string]string) error {
	values := url.Values{}
	values.Add("acl:consumerKey", c.ConsumerKey)
	for k, v := range args {
		values.Add(k, v)
	}

	url := c.BaseURL + "/" + path + "?" + values.Encode()
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
