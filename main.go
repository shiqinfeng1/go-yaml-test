package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type T struct {
	// A string
	// B struct {
	// 	RenamedC int   `yaml:"c"`
	// 	D        []int `yaml:",flow"`
	// }
	UseViper bool `yaml:"useViper"`
	Contract struct {
		Controller          string   `yaml:"Controller"`
		WhiteList           string   `yaml:"WhiteList"`
		StudentScoreStorage []string `yaml:"StudentScoreStorage"`
		StudentScoreFPSave  string   `yaml:"StudentScoreFPSave"`
		VerificationStorage []string `yaml:"VerificationStorage"`
		VerificationFPSave  string   `yaml:"VerificationFPSave"`
	}
}

func main() {

	//读取配置文件内容
	cfg, _ := filepath.Abs("./myConfig.yaml")
	input, err := ioutil.ReadFile(cfg)

	t := T{}
	err = yaml.Unmarshal(input, &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	t.Contract.StudentScoreStorage = append(t.Contract.StudentScoreStorage, "12")
	d2, err2 := yaml.Marshal(&t)
	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}
	fmt.Printf("--- t2 dump:\n%s\n\n", string(d2))
	ioutil.WriteFile("./myConfig.yaml", d2, 0777)
}
