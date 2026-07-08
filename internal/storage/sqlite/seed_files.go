package sqlite

import (
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

func (d *Database) seedFiles() ([]string, error) {

	var files []string

	err := filepath.WalkDir(
		d.options.SeedPath,
		func(path string, entry fs.DirEntry, err error) error {

			if err != nil {
				return err
			}

			if entry.IsDir() {
				return nil
			}

			if strings.HasSuffix(entry.Name(), ".sql") {
				files = append(files, path)
			}

			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	sort.Strings(files)

	return files, nil
}
