package engine

import (
	"github.com/shivakishore14/govm/utils"
	"os"
	"path/filepath"
)

func Uninstall(versionName string) error {
	config := utils.LoadConf()
	installedPath := filepath.Join(config.InstallationDir, versionName)
	if _, err := os.Stat(installedPath); err != nil {
		return err
	}
	return os.RemoveAll(installedPath)
}
