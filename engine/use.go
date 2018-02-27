package engine

import (
	"fmt"
	"github.com/shivakishore14/govm/utils"
	"log"
	"os"
	"path/filepath"
)

func Use(name string) {
	config := utils.LoadConf()
	versionPath := filepath.Join(config.InstallationDir, name, "go/bin")

	if _, err := os.Stat(versionPath); os.IsNotExist(err) {
		log.Println("Version not found")
		return
	}

	fmt.Println("PATH:", versionPath)
}
