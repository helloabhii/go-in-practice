package main

import (
	"fmt"

	"github.com/kylelemons/go-gypsy/yaml"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	config, err := yaml.ReadFile("go.yaml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Get("path"))
	fmt.Println(config.GetBool("enabled"))
}
