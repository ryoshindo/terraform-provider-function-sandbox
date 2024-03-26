package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type helloWorldDataSource struct{}

var (
	_ datasource.DataSource              = &helloWorldDataSource{}
	_ datasource.DataSourceWithConfigure = &helloWorldDataSource{}
)

func NewHelloWorldDataSource() datasource.DataSource {
	return &helloWorldDataSource{}
}

func (d *helloWorldDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_hello_world"
}

// Schema defines the schema for the data source.
func (d *helloWorldDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
			"message": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *helloWorldDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config helloWorldDataSourceModel

	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	config.Message = types.StringValue("Hello, " + config.Name.ValueString() + "!")

	diags = resp.State.Set(ctx, config)
	resp.Diagnostics.Append(diags...)
}

func (d *helloWorldDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
}

type helloWorldDataSourceModel struct {
	Name    types.String `tfsdk:"name"`
	Message types.String `tfsdk:"message"`
}
