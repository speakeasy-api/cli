package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/speakeasy-api/cli/go/cli/internal/secret"
	sdk "github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
)

// KeysClientOptions configures the keys client.
type KeysClientOptions struct {
	Scheme string
	Host   string
}

// KeysClient wraps the generated keys service client.
type KeysClient struct {
	client *sdk.Gram
}

// NewKeysClient creates a new keys client.
func NewKeysClient(options *KeysClientOptions) *KeysClient {
	return &KeysClient{client: newSDK(options.Scheme, options.Host)}
}

// NewKeysClientFromURL creates a new keys client from a URL.
func NewKeysClientFromURL(apiURL *url.URL) *KeysClient {
	return NewKeysClient(&KeysClientOptions{
		Scheme: apiURL.Scheme,
		Host:   apiURL.Host,
	})
}

// Verify validates an API key and returns organization and project info.
func (c *KeysClient) Verify(
	ctx context.Context,
	apiKey secret.Secret,
) (*shared.ValidateKeyResult, error) {
	key := apiKey.Reveal()
	result, err := c.client.Keys.Validate(ctx, operations.ValidateAPIKeySecurity{
		ApikeyHeaderGramKey: key,
	}, &key)
	if err != nil {
		return nil, fmt.Errorf("failed to verify API key: %w", err)
	}
	if result.ValidateKeyResult == nil {
		return nil, fmt.Errorf("failed to verify API key: empty response")
	}

	return result.ValidateKeyResult, nil
}
