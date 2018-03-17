package engine

import (
	"github.com/shivakishore14/govm/domain"
	"github.com/shivakishore14/govm/errors"
	"os"
	"path/filepath"
)

func Path(name string) (string, error) {
	config := domain.Config{}
	config.LoadConf()
	versionPath := filepath.Join(config.InstallationDir, name, "go")

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		return "", customErrors.ErrVersionNotFound
	}

	return versionPath, nil
}
