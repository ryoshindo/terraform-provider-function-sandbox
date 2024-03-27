package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	tffunction "github.com/ryoshindo/terraform-provider-function-sandbox/internal/function"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &fsProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &fsProvider{
			version: version,
		}
	}
}

type fsProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

type fsProviderModel struct {
	Name types.String `tfsdk:"name"`
}

// Metadata returns the provider type name.
func (p *fsProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fs"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *fsProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *fsProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config fsProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	name := ""
	if !config.Name.IsNull() {
		name = config.Name.ValueString()
	}

	if name == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("name"),
			"Name is required",
			"Name must be set to a non-empty string",
		)
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *fsProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewHelloWorldDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *fsProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *fsProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		tffunction.NewARNBuildFunction,
		tffunction.NewUlidGenerateFunction,
	}
}
