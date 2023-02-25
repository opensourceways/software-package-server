package utils

import (
	"os"
	"unicode/utf8"

	"sigs.k8s.io/yaml"
)

func LoadFromYaml(path string, cfg interface{}) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	content := []byte(os.ExpandEnv(string(b)))

	return yaml.Unmarshal(content, cfg)
}

func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}
