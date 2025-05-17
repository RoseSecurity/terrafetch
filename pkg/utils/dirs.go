package utils

import (
	"io/fs"
	"path/filepath"
	"strings"

	log "github.com/charmbracelet/log"
)

// FindTFDirs returns an array of directories containing Terraform code
func FindTFDirs(dir string) (map[string]struct{}, error) {
	tfDirs := make(map[string]struct{})

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Error("error accessing path %q: %v", path, err)
			return nil
		}

		if d.IsDir() && d.Name() == ".terraform" {
			return filepath.SkipDir
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".tf" {
			parent := filepath.Dir(path)
			tfDirs[parent] = struct{}{}
		}

		return nil
	})
	if err != nil {
		log.Error("error walking the path %q: %v", dir, err)
		return nil, err
	}

	return tfDirs, nil
}

func FindFiles(root string) (int, int, error) {
	var tfCount, docCount int

	skipDirs := map[string]bool{
		".terraform": true,
		".git":       true,
		"vendor":     true,
		"test":       true,
	}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip ignored directories
		if d.IsDir() && skipDirs[d.Name()] {
			return fs.SkipDir
		}

		name := d.Name()

		if !d.IsDir() && (strings.HasSuffix(name, ".tf") || strings.HasSuffix(name, ".tofu")) {
			tfCount++
		}

		// Match documentation files
		if !d.IsDir() &&
			(strings.HasPrefix(strings.ToLower(name), "readme") ||
				strings.HasPrefix(strings.ToLower(name), "contributing") ||
				strings.Contains(path, string(filepath.Separator)+"docs"+string(filepath.Separator)) ||
				strings.Contains(path, string(filepath.Separator)+"examples"+string(filepath.Separator))) {
			docCount++
		}

		return nil
	})

	return tfCount, docCount, err
}
