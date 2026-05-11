package api

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/speakeasy-api/gf/go/cli/internal/secret"
	sdk "github.com/speakeasy-api/gf/go/sdk"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/operations"
	"github.com/speakeasy-api/gf/go/sdk/pkg/models/shared"
)

type DeploymentsClientOptions struct {
	Scheme string
	Host   string
}

type DeploymentsClient struct {
	client *sdk.Gram
}

func NewDeploymentsClient(options *DeploymentsClientOptions) *DeploymentsClient {
	return &DeploymentsClient{client: newSDK(options.Scheme, options.Host)}
}

type CreateDeploymentRequest struct {
	APIKey          secret.Secret
	NonBlocking     bool
	ProjectSlug     string
	IdempotencyKey  string
	OpenAPIv3Assets []shared.AddOpenAPIv3DeploymentAssetForm
	Functions       []shared.AddFunctionsForm
}

func createDeploymentSecurity(key string, projectSlug string) operations.CreateDeploymentSecurity {
	return operations.CreateDeploymentSecurity{
		Option1: &operations.CreateDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func deploymentSecurity(key string, projectSlug string) operations.GetDeploymentSecurity {
	return operations.GetDeploymentSecurity{
		Option1: &operations.GetDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func latestDeploymentSecurity(key string, projectSlug string) operations.GetLatestDeploymentSecurity {
	return operations.GetLatestDeploymentSecurity{
		Option1: &operations.GetLatestDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func activeDeploymentSecurity(key string, projectSlug string) operations.GetActiveDeploymentSecurity {
	return operations.GetActiveDeploymentSecurity{
		Option1: &operations.GetActiveDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func evolveDeploymentSecurity(key string, projectSlug string) operations.EvolveDeploymentSecurity {
	return operations.EvolveDeploymentSecurity{
		Option1: &operations.EvolveDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

func redeployDeploymentSecurity(key string, projectSlug string) operations.RedeployDeploymentSecurity {
	return operations.RedeployDeploymentSecurity{
		Option1: &operations.RedeployDeploymentSecurityOption1{
			ApikeyHeaderGramKey:          key,
			ProjectSlugHeaderGramProject: projectSlug,
		},
	}
}

// CreateDeployment creates a remote deployment.
func (c *DeploymentsClient) CreateDeployment(
	ctx context.Context,
	req CreateDeploymentRequest,
) (*shared.CreateDeploymentResult, error) {
	key := req.APIKey.Reveal()
	if req.IdempotencyKey == "" {
		req.IdempotencyKey = uuid.New().String()
	}

	result, err := c.client.Deployments.Create(ctx, operations.CreateDeploymentRequest{
		GramKey:        &key,
		GramProject:    &req.ProjectSlug,
		GramSession:    nil,
		IdempotencyKey: req.IdempotencyKey,
		Body: shared.CreateDeploymentRequestBody{
			ExternalID:      nil,
			ExternalMcps:    nil,
			ExternalURL:     nil,
			Functions:       req.Functions,
			GithubPr:        nil,
			GithubRepo:      nil,
			GithubSha:       nil,
			NonBlocking:     &req.NonBlocking,
			Openapiv3Assets: req.OpenAPIv3Assets,
			Packages:        nil,
		},
	}, createDeploymentSecurity(key, req.ProjectSlug))
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.CreateDeploymentResult == nil {
		return nil, fmt.Errorf("api error: empty create deployment response")
	}

	return result.CreateDeploymentResult, nil
}

// GetDeployment retrieves a deployment by its ID.
func (c *DeploymentsClient) GetDeployment(
	ctx context.Context,
	apiKey secret.Secret,
	projectSlug string,
	deploymentID string,
) (*shared.Deployment, error) {
	key := apiKey.Reveal()
	result, err := c.client.Deployments.GetByID(
		ctx,
		deploymentSecurity(key, projectSlug),
		deploymentID,
		&key,
		nil,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.GetDeploymentResult == nil {
		return nil, fmt.Errorf("api error: empty get deployment response")
	}

	return deploymentFromGetResult(result.GetDeploymentResult), nil
}

// GetLatestDeployment retrieves the latest deployment for a project.
func (c *DeploymentsClient) GetLatestDeployment(
	ctx context.Context,
	apiKey secret.Secret,
	projectSlug string,
) (*shared.Deployment, error) {
	key := apiKey.Reveal()
	result, err := c.client.Deployments.Latest(
		ctx,
		latestDeploymentSecurity(key, projectSlug),
		&key,
		nil,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.GetLatestDeploymentResult == nil {
		return nil, fmt.Errorf("api error: empty latest deployment response")
	}

	return result.GetLatestDeploymentResult.Deployment, nil
}

// GetActiveDeployment retrieves the active deployment for a project.
func (c *DeploymentsClient) GetActiveDeployment(
	ctx context.Context,
	apiKey secret.Secret,
	projectSlug string,
) (*shared.Deployment, error) {
	key := apiKey.Reveal()
	result, err := c.client.Deployments.Active(
		ctx,
		activeDeploymentSecurity(key, projectSlug),
		&key,
		nil,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.GetActiveDeploymentResult == nil {
		return nil, fmt.Errorf("api error: empty active deployment response")
	}

	return result.GetActiveDeploymentResult.Deployment, nil
}

// EvolveRequest lists the assets to add to a deployment.
type EvolveRequest struct {
	OpenAPIv3Assets []shared.AddOpenAPIv3DeploymentAssetForm
	NonBlocking     bool
	Functions       []shared.AddFunctionsForm
	APIKey          secret.Secret
	DeploymentID    *string
	ProjectSlug     string
}

// Evolve adds assets to an existing deployment.
func (c *DeploymentsClient) Evolve(
	ctx context.Context,
	req EvolveRequest,
) (*shared.EvolveResult, error) {
	key := req.APIKey.Reveal()
	result, err := c.client.Deployments.EvolveDeployment(
		ctx,
		evolveDeploymentSecurity(key, req.ProjectSlug),
		shared.EvolveForm{
			DeploymentID:           req.DeploymentID,
			ExcludeExternalMcps:    []string{},
			ExcludeFunctions:       []string{},
			ExcludeOpenapiv3Assets: []string{},
			ExcludePackages:        []string{},
			NonBlocking:            &req.NonBlocking,
			UpsertExternalMcps:     nil,
			UpsertFunctions:        req.Functions,
			UpsertOpenapiv3Assets:  req.OpenAPIv3Assets,
			UpsertPackages:         []shared.AddPackageForm{},
		},
		&key,
		nil,
		&req.ProjectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.EvolveResult == nil {
		return nil, fmt.Errorf("api error: empty evolve response")
	}

	return result.EvolveResult, nil
}

// Redeploy triggers a redeployment of an existing deployment.
func (c *DeploymentsClient) Redeploy(
	ctx context.Context,
	apiKey secret.Secret,
	projectSlug string,
	deploymentID string,
) (*shared.Deployment, error) {
	key := apiKey.Reveal()
	result, err := c.client.Deployments.RedeployDeployment(
		ctx,
		redeployDeploymentSecurity(key, projectSlug),
		shared.RedeployRequestBody{DeploymentID: deploymentID},
		&key,
		nil,
		&projectSlug,
	)
	if err != nil {
		return nil, fmt.Errorf("api error: %w", err)
	}
	if result.RedeployResult == nil {
		return nil, fmt.Errorf("api error: empty redeploy response")
	}

	return result.RedeployResult.Deployment, nil
}

func deploymentFromGetResult(result *shared.GetDeploymentResult) *shared.Deployment {
	return &shared.Deployment{
		ClonedFrom:           result.ClonedFrom,
		CreatedAt:            result.CreatedAt,
		ExternalID:           result.ExternalID,
		ExternalMcpToolCount: result.ExternalMcpToolCount,
		ExternalMcps:         result.ExternalMcps,
		ExternalURL:          result.ExternalURL,
		FunctionsAssets:      result.FunctionsAssets,
		FunctionsToolCount:   result.FunctionsToolCount,
		GithubPr:             result.GithubPr,
		GithubRepo:           result.GithubRepo,
		GithubSha:            result.GithubSha,
		ID:                   result.ID,
		IdempotencyKey:       result.IdempotencyKey,
		Openapiv3Assets:      result.Openapiv3Assets,
		Openapiv3ToolCount:   result.Openapiv3ToolCount,
		OrganizationID:       result.OrganizationID,
		Packages:             result.Packages,
		ProjectID:            result.ProjectID,
		Status:               result.Status,
		UserID:               result.UserID,
	}
}
