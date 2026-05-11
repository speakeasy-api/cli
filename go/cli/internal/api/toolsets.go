package api

import (
	"context"
	"fmt"

	"github.com/speakeasy-api/cli/go/cli/internal/secret"
	sdk "github.com/speakeasy-api/cli/go/sdk"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/cli/go/sdk/pkg/models/shared"
)

type ToolsetsClientOptions struct {
	Scheme string
	Host   string
}

type ToolsetsClient struct {
	client *sdk.Gram
}

func NewToolsetsClient(options *ToolsetsClientOptions) *ToolsetsClient {
	return &ToolsetsClient{client: newSDK(options.Scheme, options.Host)}
}

func toolsetSecurity(key string, projectSlug string) operations.GetToolsetSecurity {
	return operations.GetToolsetSecurity{
		Option2: &operations.GetToolsetSecurityOption2{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func listToolsetsSecurity(key string, projectSlug string) operations.ListToolsetsSecurity {
	return operations.ListToolsetsSecurity{
		Option2: &operations.ListToolsetsSecurityOption2{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func (c *ToolsetsClient) GetToolset(
	ctx context.Context,
	apiKey secret.Secret,
	projectSlug string,
	toolsetSlug string,
) (*shared.Toolset, error) {
	key := apiKey.Reveal()
	result, err := c.client.Toolsets.GetBySlug(
		ctx,
		toolsetSecurity(key, projectSlug),
		toolsetSlug,
		nil,
		&key,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get toolset: %w", err)
	}
	if result.Toolset == nil {
		return nil, fmt.Errorf("failed to get toolset: empty response")
	}

	return result.Toolset, nil
}

func (c *ToolsetsClient) ListToolsets(ctx context.Context, apiKey secret.Secret, projectSlug string) ([]shared.ToolsetEntry, error) {
	key := apiKey.Reveal()
	result, err := c.client.Toolsets.List(
		ctx,
		listToolsetsSecurity(key, projectSlug),
		nil,
		&key,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list toolsets: %w", err)
	}
	if result.ListToolsetsResult == nil {
		return nil, fmt.Errorf("failed to list toolsets: empty response")
	}

	return result.ListToolsetsResult.Toolsets, nil
}
