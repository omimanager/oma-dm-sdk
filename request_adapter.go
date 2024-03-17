package omadm

import (
	"context"
)

// RequestAdapter defines the interface for sending requests and handling responses.
type RequestAdapter interface {
	SendRequest(ctx context.Context, endpointURL string, credentials map[string]string, requestData []byte) ([]byte, error)
}
