package internal

import (
	"github.com/RoseSecurity/terrafetch/pkg/utils"
	log "github.com/charmbracelet/log"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type Analytics struct {
	VariableCount          int
	SensitiveVariableCount int
	ResourceCount          int
	OutputCount            int
	SensitiveOutputCount   int
	DataSourceCount        int
	ProviderCount          int
	ModuleCount            int
	FileCount              int
	DocCount               int
}

func AnalyzeRepository(rootDir string) ([]Analytics, error) {
	dirs, err := utils.FindTFDirs(rootDir)
	if err != nil {
		return nil, ErrFailedToFindDir
	}

	if len(dirs) == 0 {
		return nil, ErrNoTerraformFiles
	}

	var totalVars, totalResources, totalOutputs, totalDataSources, totalModules, totalProviders, totalSensitiveVars, totalSensitiveOutputs int

	for dir := range dirs {
		if !isTerraformDirectory(dir) {
			continue
		}

		repo, diags := tfconfig.LoadModule(dir)
		if diags.HasErrors() {
			log.Warn("could not load %v", dir)
		}

		totalVars += len(repo.Variables)
		totalOutputs += len(repo.Outputs)
		totalResources += len(repo.ManagedResources)
		totalDataSources += len(repo.DataResources)
		totalModules += len(repo.ModuleCalls)
		totalProviders += len(repo.RequiredProviders)

		for _, v := range repo.Variables {
			if v.Sensitive {
				totalSensitiveVars++
			}
		}

		for _, v := range repo.Outputs {
			if v.Sensitive {
				totalSensitiveOutputs++
			}
		}
	}

	totalTfFiles, totalDocFiles, err := utils.FindFiles(rootDir)
	if err != nil {
		log.Error("could not count terraform files %v", err)
	}

	return []Analytics{
		{
			VariableCount:          totalVars,
			SensitiveVariableCount: totalSensitiveVars,
			ResourceCount:          totalResources,
			OutputCount:            totalOutputs,
			SensitiveOutputCount:   totalSensitiveOutputs,
			DataSourceCount:        totalDataSources,
			ProviderCount:          totalProviders,
			ModuleCount:            totalModules,
			FileCount:              totalTfFiles,
			DocCount:               totalDocFiles,
		},
	}, nil
}

// isTerraformDirectory returns if a directory contains Terraform code
func isTerraformDirectory(dir string) bool {
	return tfconfig.IsModuleDir(dir)
}
