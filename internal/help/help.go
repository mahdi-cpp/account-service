package help

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

func ToStringJson[T any](data T) (string, error) {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}
	return string(jsonBytes), nil
}

// GetUserId retrieves the userID from the request header.
// It returns the userID string and an error if the header is missing.
func GetUserId(c *gin.Context) (string, error) {
	// Get the "userID" header from the request.
	// If the header is not present, c.GetHeader will return an empty string "".
	userIDStr := c.GetHeader("userID")

	// Check if the retrieved userID string is empty.
	if userIDStr == "" {
		// If it's empty, return an empty string and a new error.
		// errors.New is used for simple, static error messages.
		// fmt.Errorf is useful if you need to include dynamic information in your error message.
		return "", errors.New("userID header is missing or empty")
		// Or using fmt.Errorf for a more descriptive error:
		// return "", fmt.Errorf("missing or empty 'userID' header in request for path: %s", c.Request.URL.Path)
	}

	// If the userID is found and not empty, return it and a nil error.
	return userIDStr, nil
}

// MakeRequest Helper function to make HTTP requests
func MakeRequest(t *testing.T, method, endpoint string, queryParams map[string]interface{}, body interface{}) ([]byte, error) {

	// Build URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	if queryParams != nil {
		q := u.Query()
		for key, value := range queryParams {
			q.Add(key, fmt.Sprintf("%v", value))
		}
		u.RawQuery = q.Encode()
	}

	// Marshal body if provided
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshaling body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	fmt.Println(u.String())
	fmt.Println("")

	// create request
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, resp.Status)
	}

	// ReadChat response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	return respBody, nil
}

func MakeRequestBody(t *testing.T, method, endpoint string, body interface{}) (*http.Response, error) {

	// Build URL with query parameters
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %w", err)
	}

	// Marshal body if provided
	var bodyReader io.Reader

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("marshaling body: %w", err)
	}
	bodyReader = bytes.NewReader(jsonData)

	fmt.Println(u.String())
	fmt.Println("")

	// create request
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}

	return resp, nil
}

func StrPtr(str string) *string {
	return &str
}

func BoolPtr(bool bool) *bool {
	return &bool
}
