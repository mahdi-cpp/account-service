package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/goccy/go-json"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/mahdi-cpp/account-service/internal/collections/user"
	"github.com/mahdi-cpp/account-service/internal/help"
)

const baseURL = "http://localhost:50002/api/"

func TestCreateUser(t *testing.T) {

	currentURL := baseURL + "users"

	body := &user.User{
		Username:    "parsa",
		Email:       "parsa@gmail.com",
		PhoneNumber: "+989123000200",
		FirstName:   "Mahdi",
		LastName:    "Abdolmaleki",
		IsVerified:  true,
	}

	resp, err := help.MakeRequestBody(t, "POST", currentURL, body)
	if err != nil {
		t.Fatalf("create request failed: %v", err) // از t.Fatalf به جای t.Errorf استفاده کنید تا تست بلافاصله متوقف شود
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		var r Error
		if err := json.Unmarshal(respBody, &r); err != nil {
			t.Fatalf("unmarshaling response: %v", err)
		}
		t.Fatalf("error %s", r.Message)
	}

	var createUser user.User
	if err := json.Unmarshal(respBody, &createUser); err != nil { // از اشاره‌گر (&) استفاده کنید
		t.Fatalf("unmarshaling response: %v", err)
	}

	t.Logf("Created user ID: %s", createUser.ID)

	if diff := cmp.Diff(body.Username, createUser.Username); diff != "" {
		t.Errorf("Caption mismatch (-want +got):\n%s", diff)
	}
}

func TestReadUser(t *testing.T) {

}

func TestReadUsers(t *testing.T) {

}

func TestUpdateUser(t *testing.T) {

	currentURL := baseURL + "users"
	id, err := uuid.Parse("0199778a-8e7a-7dfc-8aa0-e41c8d53fc74")
	if err != nil {
		t.Fatalf("parse uuid failed: %v", err)
	}

	body := &user.UpdateOptions{
		ID: id,
		//Username: help.StrPtr("parsa"),
		//Email: help.StrPtr("parsa@gmail.com"),
		//PhoneNumber: help.StrPtr("+98912300015"),
		FirstName:  help.StrPtr("Mahdi"),
		LastName:   help.StrPtr("Abdolmaleki"),
		IsVerified: help.BoolPtr(true),
	}

	resp, err := help.MakeRequestBody(t, "PATCH", currentURL, body)
	if err != nil {
		t.Fatalf("update request failed: %v", err) // از t.Fatalf به جای t.Errorf استفاده کنید تا تست بلافاصله متوقف شود
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("reading response: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		var r Error
		if err := json.Unmarshal(respBody, &r); err != nil {
			t.Fatalf("unmarshaling response: %v", err)
		}
		t.Fatalf("error %s", r.Message)
	}

	var createUser user.User
	if err := json.Unmarshal(respBody, &createUser); err != nil { // از اشاره‌گر (&) استفاده کنید
		t.Fatalf("unmarshaling response: %v", err)
	}

	t.Logf("update user ID: %s", createUser.ID)

}

func TestDeleteUser(t *testing.T) {}
