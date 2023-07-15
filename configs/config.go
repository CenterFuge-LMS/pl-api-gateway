package configs

import (
	"log"
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

func NewConfig() *koanf.Koanf {
	k := koanf.New(".")
	err := k.Load(file.Provider(configPath()), yaml.Parser())
	if err != nil {
		log.Fatal("error loading config", err)

	}

	return k
}
func configPath() string {
	filePath := []string{

		"configs/dev.yaml",
	}
	for _, file := range filePath {
		if _, err := os.Stat(file); err == nil {
			return file

		}

	}
	return ""
}
