package engine

import (
	"github.com/shivakishore14/govm/domain"
	"io/ioutil"
	"log"
	"strings"
)

func LocalList() domain.Versions {
	config := domain.Config{}
	config.LoadConf()
	files, err := ioutil.ReadDir(config.InstallationDir)
	if err != nil {
		log.Fatal(err)
	}
	versions := domain.Versions{}
	for _, f := range files {
		if f.IsDir() {
			name := f.Name()
			if strings.HasPrefix(name, "go") {
				version := domain.Version{}
				version.Name = name
				versions = append(versions, version)
			}
		}
	}
	return versions
}
