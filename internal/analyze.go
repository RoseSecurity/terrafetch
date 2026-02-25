package internal

import (
	"runtime"
	"sync"

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
	scan, err := utils.ScanRepository(rootDir)
	if err != nil {
		return nil, ErrFailedToFindDir
	}

	if len(scan.TFDirs) == 0 {
		return nil, ErrNoTerraformFiles
	}

	// Parallelize module analysis with a worker pool
	var wg sync.WaitGroup
	results := make(chan Analytics, len(scan.TFDirs))
	sem := make(chan struct{}, runtime.NumCPU())

	for dir := range scan.TFDirs {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			if !isTerraformDirectory(d) {
				return
			}

			repo, diags := tfconfig.LoadModule(d)
			if diags.HasErrors() {
				log.Warn("could not load %v", d)
				return
			}

			var a Analytics
			a.VariableCount = len(repo.Variables)
			a.OutputCount = len(repo.Outputs)
			a.ResourceCount = len(repo.ManagedResources)
			a.DataSourceCount = len(repo.DataResources)
			a.ModuleCount = len(repo.ModuleCalls)
			a.ProviderCount = len(repo.RequiredProviders)

			for _, v := range repo.Variables {
				if v.Sensitive {
					a.SensitiveVariableCount++
				}
			}

			for _, v := range repo.Outputs {
				if v.Sensitive {
					a.SensitiveOutputCount++
				}
			}

			results <- a
		}(dir)
	}

	// Close results channel after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Aggregate results from all workers
	var total Analytics
	for a := range results {
		total.VariableCount += a.VariableCount
		total.SensitiveVariableCount += a.SensitiveVariableCount
		total.ResourceCount += a.ResourceCount
		total.OutputCount += a.OutputCount
		total.SensitiveOutputCount += a.SensitiveOutputCount
		total.DataSourceCount += a.DataSourceCount
		total.ProviderCount += a.ProviderCount
		total.ModuleCount += a.ModuleCount
	}

	total.FileCount = scan.TFCount
	total.DocCount = scan.DocCount

	return []Analytics{total}, nil
}

// isTerraformDirectory returns if a directory contains Terraform code
func isTerraformDirectory(dir string) bool {
	return tfconfig.IsModuleDir(dir)
}
