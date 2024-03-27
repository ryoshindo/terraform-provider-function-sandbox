package function

import (
	"context"
	"math/rand"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/oklog/ulid/v2"
)

var _ function.Function = ulidGenerateFunction{}

func NewUlidGenerateFunction() function.Function {
	return &ulidGenerateFunction{}
}

type ulidGenerateFunction struct{}

func (f ulidGenerateFunction) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "ulid_generate"
}

func (f ulidGenerateFunction) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "ulid_generate Function",
		MarkdownDescription: "Generate a ULID",
		Return:              function.StringReturn{},
	}
}

func (f ulidGenerateFunction) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)

	if err != nil {
		return
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, id.String()))
}
