package internal

import (
	"github.com/RoseSecurity/terrafetch/pkg/utils"
	log "github.com/charmbracelet/log"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type Analytics struct {
	VariableCount   int
	ResourceCount   int
	OutputCount     int
	DataSourceCount int
	ProviderCount   int
	ModuleCount     int
}

func AnalyzeRepository(rootDir string) ([]Analytics, error) {
	dirs, err := utils.FindTFDirs(rootDir)
	if err != nil {
		return nil, ErrFailedToFindDir
	}

	var totalVars, totalResources, totalOutputs, totalDataSources, totalModules, totalProviders int

	for dir := range dirs {
		if !isTerraformDirectory(dir) {
			continue
		}

		repo, diags := tfconfig.LoadModule(dir)
		if diags.HasErrors() {
			log.Warn("could not load %d", dir)
		}

		totalVars += len(repo.Variables)
		totalOutputs += len(repo.Outputs)
		totalResources += len(repo.ManagedResources)
		totalDataSources += len(repo.DataResources)
		totalModules += len(repo.ModuleCalls)
		totalProviders += len(repo.RequiredProviders)
	}

	return []Analytics{
		{
			VariableCount:   totalVars,
			ResourceCount:   totalResources,
			OutputCount:     totalOutputs,
			DataSourceCount: totalDataSources,
			ProviderCount:   totalProviders,
			ModuleCount:     totalModules,
		},
	}, nil
}

// isTerraformDirectory returns if a directory contains Terraform code
func isTerraformDirectory(dir string) bool {
	return tfconfig.IsModuleDir(dir)
}
