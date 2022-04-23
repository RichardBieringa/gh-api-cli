package api

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestGetUser(t *testing.T) {
	content, err := os.ReadFile("./test-responses/user.json")
	if err != nil {
		t.Fatal("could not read file")
	}

	var user User
	err = json.Unmarshal(content, &user)
	if err != nil {
		t.Fatal("could not decode users test response")
	}

	got, err := GetUser("RichardBieringa")
	t.Run("the request should not have an error", func(t *testing.T) {
		if err != nil {
			t.Fatalf("received an error %v", err)
		}
	})

	t.Run("the response should contain the user response for richardbieringa", func(t *testing.T) {
		if !reflect.DeepEqual(got, user) {
			t.Errorf("Unexpected Response, got %v, want %v", got, user)
		}
	})
}

func TestGetUsers(t *testing.T) {
	content, err := os.ReadFile("./test-responses/users.json")
	if err != nil {
		t.Fatal("could not read file")
	}

	users := make([]User, 0)
	err = json.Unmarshal(content, &users)
	if err != nil {
		t.Fatal("could not decode users test response")
	}

	got, err := GetUsers()
	t.Run("the request should not have an error", func(t *testing.T) {
		if err != nil {
			t.Fatalf("received an error %v", err)
		}
	})

	t.Run("the response should contain the first 30 users", func(t *testing.T) {
		if !reflect.DeepEqual(got, users) {
			t.Errorf("Unexpected Response, got %v, want %v", got, users)
		}
	})
}
