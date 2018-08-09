package oauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/battlenet"
)

type provider struct {
	config  *oauth2.Config
	getUser func(*http.Client) (*user, error)
}

type providerConfig struct {
	endpoint oauth2.Endpoint
	scopes   []string
	getUser  func(*http.Client) (*user, error)
}

type user struct {
	id   string
	name string
}

var providerConfigs = map[string]providerConfig{
	"google": {
		endpoint: google.Endpoint,
		scopes:   []string{"profile"},
		getUser:  getGoogleUser,
	},
	"facebook": {
		endpoint: facebook.Endpoint,
		scopes:   []string{"public_profile","email"},
		getUser:  getFacebookUser,
	},
	"github": {
		endpoint: github.Endpoint,
		scopes:   []string{},
		getUser:  getGithubUser,
	},
	"battlenet": {
		endpoint: battlenet.Endpoint,
		scopes:   []string{"sc2.profile"},
		getUser:  getBattlenetUser,
	},
}

func getGoogleUser(c *http.Client) (*user, error) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"

	u := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{}

	err := getJSON(c, url, &u)
	if err != nil {
		return nil, err
	}

	return &user{id: u.ID, name: u.Name}, nil
}

func getFacebookUser(c *http.Client) (*user, error) {
	url := "https://graph.facebook.com/me?fields=id,name"

	u := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{}

	err := getJSON(c, url, &u)
	if err != nil {
		return nil, err
	}

	return &user{id: u.ID, name: u.Name}, nil
}

func getGithubUser(c *http.Client) (*user, error) {
	url := "https://api.github.com/user"

	u := struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}{}

	err := getJSON(c, url, &u)
	if err != nil {
		return nil, err
	}

	return &user{id: strconv.FormatInt(u.ID, 10), name: u.Name}, nil
}

/*
type AccountService struct {
	client *Client
}

// User represents the user information for a Battle.net account
type User struct {
	ID        int    `json:"id"`
	BattleTag string `json:"battletag"`
}

// User calls the /account/user endpoint. See Battle.net docs.
func (s *AccountService) User() (*User, *Response, error) {
	req, err := s.client.NewRequest("GET", "account/user", nil)
	if err != nil {
		return nil, nil, err
	}

	var user User
	resp, err := s.client.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return &user, resp, nil
}*/
func getBattlenetUser(c *http.Client) (*user, error) {
	url := "https://api.battlenet.com.cn/account/user"

	u := struct {
		ID   int64  `json:"id"`
		Name string `json:"battletag"`
	}{}

	err := getJSON(c, url, &u)
	if err != nil {
		return nil, err
	}

	return &user{id: strconv.FormatInt(u.ID, 10), name: u.Name}, nil
}

func getJSON(c *http.Client, url string, v interface{}) error {
	response, err := c.Get(url)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode > 299 {
		return fmt.Errorf("bad request status code: %v", response.StatusCode)
	}

	err = json.NewDecoder(io.LimitReader(response.Body, 1<<20)).Decode(v)
	if err != nil {
		return fmt.Errorf("failed to decode JSON response: %v", err)
	}

	return nil
}
