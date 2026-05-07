package api

import (
	"context"
	"fmt"
	"io"

	sdk "github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/shared"
)

type AssetsClientOptions struct {
	Scheme string
	Host   string
}

type AssetsClient struct {
	client *sdk.Gram
}

func NewAssetsClient(options *AssetsClientOptions) *AssetsClient {
	return &AssetsClient{client: newSDK(options.Scheme, options.Host)}
}

type UploadAssetForm struct {
	APIKey        string
	ProjectSlug   string
	Reader        io.ReadCloser
	ContentType   string
	ContentLength int64
}

func (c *AssetsClient) UploadOpenAPIv3(
	ctx context.Context,
	req *UploadAssetForm,
) (*shared.Asset, error) {
	result, err := c.client.Assets.UploadOpenAPIv3(
		ctx,
		operations.UploadOpenAPIv3AssetRequest{
			ContentLength: req.ContentLength,
			GramKey:       &req.APIKey,
			GramProject:   &req.ProjectSlug,
			GramSession:   nil,
			Body:          req.Reader,
		},
		operations.UploadOpenAPIv3AssetSecurity{
			Option1: &operations.UploadOpenAPIv3AssetSecurityOption1{
				ApikeyHeaderGramKey:          req.APIKey,
				ProjectSlugHeaderGramProject: req.ProjectSlug,
			},
		},
		operations.WithSetHeaders(map[string]string{"Content-Type": req.ContentType}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload OpenAPI source: %w", err)
	}
	if result.UploadOpenAPIv3Result == nil {
		return nil, fmt.Errorf("failed to upload OpenAPI source: empty response")
	}

	return &result.UploadOpenAPIv3Result.Asset, nil
}

func (c *AssetsClient) UploadFunctions(
	ctx context.Context,
	req *UploadAssetForm,
) (*shared.Asset, error) {
	result, err := c.client.Assets.UploadFunctions(
		ctx,
		operations.UploadFunctionsRequest{
			ContentLength: req.ContentLength,
			GramKey:       &req.APIKey,
			GramProject:   &req.ProjectSlug,
			GramSession:   nil,
			Body:          req.Reader,
		},
		operations.UploadFunctionsSecurity{
			Option1: &operations.UploadFunctionsSecurityOption1{
				ApikeyHeaderGramKey:          req.APIKey,
				ProjectSlugHeaderGramProject: req.ProjectSlug,
			},
		},
		operations.WithSetHeaders(map[string]string{"Content-Type": req.ContentType}),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to upload Functions source: %w", err)
	}
	if result.UploadFunctionsResult == nil {
		return nil, fmt.Errorf("failed to upload Functions source: empty response")
	}

	return &result.UploadFunctionsResult.Asset, nil
}
