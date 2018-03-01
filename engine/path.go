package engine

import (
	"github.com/shivakishore14/govm/utils"
	"log"
	"os"
	"path/filepath"
)

func Path(name string) string{
	config := utils.LoadConf()
	versionPath := filepath.Join(config.InstallationDir, name, "go/bin")

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		log.Println("Version not found")
		return ""
	}

	return "PATH:" + versionPath
}