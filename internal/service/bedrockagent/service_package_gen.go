// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package bedrockagent

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	bedrockagent_sdkv2 "github.com/aws/aws-sdk-go-v2/service/bedrockagent"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceAgentVersions,
			Name:    "Agent Versions",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newAgentActionGroupResource,
			Name:    "Agent Action Group",
		},
		{
			Factory: newAgentAliasResource,
			Name:    "Agent Alias",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "agent_alias_arn",
			},
		},
		{
			Factory: newAgentKnowledgeBaseAssociationResource,
			Name:    "Agent Knowledge Base Association",
		},
		{
			Factory: newAgentResource,
			Name:    "Agent",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "agent_arn",
			},
		},
		{
			Factory: newDataSourceResource,
			Name:    "Data Source",
		},
		{
			Factory: newKnowledgeBaseResource,
			Name:    "Knowledge Base",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.BedrockAgent
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*bedrockagent_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return bedrockagent_sdkv2.NewFromConfig(cfg,
		bedrockagent_sdkv2.WithEndpointResolverV2(newEndpointResolverSDKv2()),
		withBaseEndpoint(config[names.AttrEndpoint].(string)),
	), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}