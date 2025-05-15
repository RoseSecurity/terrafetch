package internal

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type Analytics struct {
	VariableCount int
	ResourceCount int
	OutputCount   int
}

func AnalyzeRepository(dir string) ([]Analytics, error) {
	if !isTerraformDirectory(dir) {
		return nil, ErrDirMissingCode
	}

	repo, diags := tfconfig.LoadModule(dir)
	if diags.HasErrors() {
		return nil, diags
	}

	return []Analytics{
		{
			VariableCount: len(repo.Variables),
			ResourceCount: len(repo.ManagedResources),
			OutputCount:   len(repo.Outputs),
		},
	}, nil
}

func isTerraformDirectory(dir string) bool {
	return tfconfig.IsModuleDir(dir)
}
