package terraform

import (
	"fmt"
	"github.com/thazelart/terraform-validator/internal/utils"
	"strings"
)

func RunTerraformFmt(filepath string) (string, bool) {
	utils.EnsureProgramInstalled("terraform")

	stdout, _ := utils.RunSystemCommand("terraform", "fmt", "-write=false", filepath)

	if stdout != "" {
		fmt.Println("\nERROR: Non formatted file(s):")
		for _, file := range strings.Split(stdout, "\n") {
			fmt.Printf("  - %s\n", file)
		}
		return stdout, false
	}
	return "", true
}
