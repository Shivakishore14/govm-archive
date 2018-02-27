package utils

import "github.com/shivakishore14/govm/domain"

func LoadConf() domain.Config {
	return domain.Config{TempDir: "/Users/shiva/.govm/tmp/", InstallationDir: "/Users/shiva/.govm/installed/"}
}
