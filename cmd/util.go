/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

func getOss() oss {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MAC operating system")
		return &macos{
			name: "macos",
		}
	case "linux":
		fmt.Println("Linux")
		return &linux{
			name: "linux",
		}
	default:
		log.Println("OS not supported")
		return nil
	}
}

type data struct {
	Exe          string   `yaml:"exe"`
	Config       []string `yaml:"config"`
	Dependencies []string `yaml:"dependencies"`
}

func getYaml() map[string]data {
	obj := make(map[string]data)
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, obj)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return obj
}
