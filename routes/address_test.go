package routes

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAddressesHandler(t *testing.T) {
	// Create a new request
	req, err := http.NewRequest(http.MethodGet, "http://localhost:1323/api/v1/addresses", nil)
	if err != nil {
		t.Fatal(err)
	}

	// add basic auth authentication headers
	req.SetBasicAuth("test_admin", "Password")

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// Check the response status code
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Check the response body
	expectedBody := "[{\"ID\":1,\"CreatedAt\":\"2024-01-14T21:32:13.735713645+01:00\",\"UpdatedAt\":\"2024-01-14T21:32:13.735713645+01:00\",\"DeletedAt\":null,\"street\":\"Musterstrasse\",\"postal\":\"12345\",\"city\":\"Musterhausen\",\"number\":\"1a\",\"owner_id\":0,\"owner_type\":\"\"},{\"ID\":2,\"CreatedAt\":\"2024-01-14T21:35:02.607390209+01:00\",\"UpdatedAt\":\"2024-01-14T21:35:02.607390209+01:00\",\"DeletedAt\":null,\"street\":\"Musterweg\",\"postal\":\"12345\",\"city\":\"Musterhausen\",\"number\":\"3\",\"owner_id\":1,\"owner_type\":\"students\"},{\"ID\":3,\"CreatedAt\":\"2024-01-14T21:35:08.380673281+01:00\",\"UpdatedAt\":\"2024-01-14T21:35:08.380673281+01:00\",\"DeletedAt\":null,\"street\":\"Musterweg\",\"postal\":\"12345\",\"city\":\"Musterhausen\",\"number\":\"3\",\"owner_id\":2,\"owner_type\":\"students\"},{\"ID\":4,\"CreatedAt\":\"2024-01-14T21:35:12.350381022+01:00\",\"UpdatedAt\":\"2024-01-14T21:35:12.350381022+01:00\",\"DeletedAt\":null,\"street\":\"Musterweg\",\"postal\":\"12345\",\"city\":\"Musterhausen\",\"number\":\"3\",\"owner_id\":3,\"owner_type\":\"students\"},{\"ID\":5,\"CreatedAt\":\"2024-01-14T21:35:16.688904788+01:00\",\"UpdatedAt\":\"2024-01-14T21:35:16.688904788+01:00\",\"DeletedAt\":null,\"street\":\"Musterweg\",\"postal\":\"12345\",\"city\":\"Musterhausen\",\"number\":\"3\",\"owner_id\":4,\"owner_type\":\"students\"},{\"ID\":6,\"CreatedAt\":\"2024-01-14T22:51:52.755186357+01:00\",\"UpdatedAt\":\"2024-01-14T22:51:52.755186357+01:00\",\"DeletedAt\":null,\"street\":\"Main St\",\"postal\":\"12345\",\"city\":\"San Francisco\",\"number\":\"123\",\"owner_id\":0,\"owner_type\":\"\"}]\n"
	//expectedBody := `[{"street":"Main St","city":"San Francisco", "number": "123", "postal":"12345"}]`
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedBody, string(body))
}
