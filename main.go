package main

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/config"
)

const (
	version = "0.0.1"
)

func main() {
	globalConfig := config.GenerateGlobalConfig(version)

	fmt.Printf("%v", globalConfig.WorkDir.Path)
}
