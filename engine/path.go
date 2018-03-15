package engine

import (
	"github.com/shivakishore14/govm/domain"
	"log"
	"os"
	"path/filepath"
)

func Path(name string) string {
	config := domain.Config{}
	config.LoadConf()
	versionPath := filepath.Join(config.InstallationDir, name, "go")

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		log.Println("Version not found")
		return ""
	}

	return "PATH:" + versionPath
}
