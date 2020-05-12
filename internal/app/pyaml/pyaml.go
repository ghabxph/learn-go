package pyaml

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

type Pyaml struct {
	Pyaml []struct {
		Step    string `yaml:"step"`
		WorkDir string `yaml:"workdir"`
		IfExist struct {
			Path   string     `yaml:"path"`
			Script [][]string `yaml:"script"`
		} `yaml:"ifexist"`
		Script [][]string `yaml:"script"`
	} `yaml:"pyaml"`
}

func Execute(yamlFile string) {
	var pyaml Pyaml
	yaml.Unmarshal(r(yamlFile, &pyaml))
	for _, Pyaml := range pyaml.Pyaml {
		if Pyaml.Step != "" {
			log.Println(Pyaml.Step)
			log.Println("----------------------------")
		}
		if Pyaml.WorkDir != "" {
			wd, _ := os.Getwd()
			log.Printf("Current work directory: %v\n", wd)
			log.Printf("Setting work directory to: %v\n", Pyaml.WorkDir)
			err := os.Chdir(Pyaml.WorkDir)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			wd, _ = os.Getwd()
			log.Printf("Current work directory is now on: %v\n", wd)
		}
		if len(Pyaml.Script) > 0 {
			executeScript(Pyaml.Script)
		}
		log.Println("")
		log.Println("")

	}
}

func r(yamlFile string, pyaml *Pyaml) ([]byte, **Pyaml) {
	f, _ := ioutil.ReadFile(yamlFile)
	return f, &pyaml
}

func executeScript(cmds [][]string) {
	log.Println("Executing script: ")
	for _, cmd := range cmds {
		log.Printf("- %v\n", cmd)
		out, err := exec.Command(cmd[0], cmd[1:]...).Output()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(out))
	}
}
