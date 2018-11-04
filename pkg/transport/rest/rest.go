package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/earies/ycclient/pkg/schema"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HTTPClient *http.Client
}

type Credentials struct {
	Username string
	Password string
}

func (c *Client) Contributors() ([]string, error) {
	rel := &url.URL{Path: "/api/contributors"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result schema.ContributorsResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	sort.Strings(result.Contributors)
	return result.Contributors, nil
}

// Take in generic interface or ...
// In addition to client
// Return unmarshaled payload (possibly another interface)
func restHandler(c *Client) {
}

// Can the reused code be generalized a bit more into global func
func (c *Client) Module(keys schema.ModelKeys, filter schema.SearchFilter) (schema.ModuleResult, error) {
	var result schema.ModuleResult

	modulePath := fmt.Sprintf("%s,%s,%s", keys.Name, keys.Revision, keys.Organization)
	rel := &url.URL{Path: "/api/search/modules/" + modulePath}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return result, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
