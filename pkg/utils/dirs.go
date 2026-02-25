package utils

import (
	"io/fs"
	"path/filepath"
	"strings"

	log "github.com/charmbracelet/log"
)

// ScanResult holds the results of a single filesystem walk.
type ScanResult struct {
	TFDirs  map[string]struct{}
	TFCount int
	DocCount int
}

// skipDirs contains directories that should never be descended into.
var skipDirs = map[string]bool{
	".terraform":       true,
	".terragrunt-cache": true,
	".git":             true,
	"vendor":           true,
	"test":             true,
	"node_modules":     true,
}

// ScanRepository walks the directory tree once, collecting Terraform module
// directories, file counts, and documentation file counts in a single pass.
func ScanRepository(root string) (*ScanResult, error) {
	result := &ScanResult{
		TFDirs: make(map[string]struct{}),
	}

	sep := string(filepath.Separator)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Error("error accessing path %q: %v", path, err)
			return nil
		}

		// Skip irrelevant directories as early as possible
		if d.IsDir() && skipDirs[d.Name()] {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		name := d.Name()
		ext := filepath.Ext(name)

		// Count Terraform files and track their parent directories
		if ext == ".tf" || ext == ".tofu" {
			result.TFCount++
			parent := filepath.Dir(path)
			result.TFDirs[parent] = struct{}{}
		}

		// Count documentation files
		lower := strings.ToLower(name)
		if strings.HasPrefix(lower, "readme") ||
			strings.HasPrefix(lower, "contributing") ||
			strings.Contains(path, sep+"docs"+sep) ||
			strings.Contains(path, sep+"examples"+sep) {
			result.DocCount++
		}

		return nil
	})
	if err != nil {
		log.Error("error walking the path %q: %v", root, err)
		return nil, err
	}

	return result, nil
}
