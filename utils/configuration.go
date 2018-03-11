package utils

import (
	"github.com/shivakishore14/govm/domain"
	"log"
	"os/user"
	"path/filepath"
)

func LoadConf() domain.Config {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	userHome := usr.HomeDir
	govmHome := filepath.Join(userHome, ".govm/")
	tempDir := filepath.Join(govmHome, "tmp/")
	installationDir := filepath.Join(govmHome, "installed/")
	return domain.Config{TempDir: tempDir, InstallationDir: installationDir}
}
