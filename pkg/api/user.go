package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Represents a single user in the GitHub REST API
type User struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Gets a single user by username.
func GetUser(username string) (User, error) {
	var user User

	// Make the request
	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s", username))
	if err != nil {
		return user, err
	}

	// Read the response content
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return user, fmt.Errorf("got status %v", res.StatusCode)
	}

	if err != nil {
		return user, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// Gets the GitHub users in incremental ID order.
func GetUsers() ([]User, error) {
	users := make([]User, 30)

	// Make the request
	res, err := http.Get("https://api.github.com/users")
	if err != nil {
		return users, err
	}

	// Read the response content
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return users, fmt.Errorf("got status %v", res.StatusCode)
	}

	if err != nil {
		return users, err
	}

	err = json.Unmarshal(body, &users)
	if err != nil {
		return users, err
	}

	return users, err
}
