package utils

import (
	"io/fs"
	"path/filepath"

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
