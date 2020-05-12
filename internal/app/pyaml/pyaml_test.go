package pyaml

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

type test struct {
	Test []struct {
		Title string `yaml:"title"`
		Path  string `yaml:"path"`
	} `yaml:"test"`
}

func TestExecute(t *testing.T) {
	// Sets the test environment
	dir, _ := os.Getwd()
	dir, _ = filepath.Abs(dir + "/../../../tests/internal/app/pyaml")
	tf := dir + "/pyaml_test.yaml"
	var tt test

	// Initializes test table (read yaml)
	f, _ := ioutil.ReadFile(tf)
	yaml.Unmarshal(f, &tt)

	for _, tc := range tt.Test {
		t.Run(tc.Title, func(t *testing.T) {
			Execute(dir + "/" + tc.Path)
		})
	}
}
