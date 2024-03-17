package omadm

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPRequestAdapter implements the RequestAdapter interface using HTTP.
type HTTPRequestAdapter struct {
	client *http.Client
	authProvider AuthenticationProvider
}

// NewHTTPRequestAdapter creates a new HTTPRequestAdapter.
func NewHTTPRequestAdapter(authProvider AuthenticationProvider) *HTTPRequestAdapter {
	return &HTTPRequestAdapter{
		client: &http.Client{},
		authProvider: AuthenticationProvider,
	}
}

// SendRequest sends an HTTP request to the specified endpoint.
func (a *HTTPRequestAdapter) SendRequest(ctx context.Context, requestInfo RequestInformation) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpointURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Set authentication headers (e.g., bearer token, API key)
	for key, value := range credentials {
		req.Header.Set(key, value)
	}

	// Set request data (e.g., OMA DM sync message XML)
	req.Body = ioutil.NopCloser(bytes.NewReader(requestData))
	req.Header.Set("Content-Type", "application/xml")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return responseData, nil
}