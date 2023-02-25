package utils

import (
	"io/ioutil"
	"os"
	"time"
	"unicode/utf8"

	"sigs.k8s.io/yaml"
)

func LoadFromYaml(path string, cfg interface{}) error {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	content := []byte(os.ExpandEnv(string(b)))

	return yaml.Unmarshal(content, cfg)
}

func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}

func Now() int64 {
	return time.Now().Unix()
}