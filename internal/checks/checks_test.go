package checks_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/kami-zh/go-capturer"
	"github.com/thazelart/terraform-validator/internal/checks"
	"github.com/thazelart/terraform-validator/internal/hcl"
	"testing"
)

var parsedFile = hcl.ParsedFile{
	Name: "main.tf",
	Blocks: hcl.TerraformBlocks{
		Variables: []hcl.Variable{
			hcl.Variable{Name: "var_with_description", Description: "a var description"},
			hcl.Variable{Name: "var_without_description", Description: ""},
		},
		Outputs: []hcl.Output{
			hcl.Output{Name: "out_with_description", Description: "a output description"},
			hcl.Output{Name: "out_without_description", Description: ""},
		},
		Resources: []hcl.Resource{
			hcl.Resource{Name: "a_resource", Type: "google_sql_database"},
		},
		Locals: []hcl.Locals{
			hcl.Locals{"a_local", "another_local"},
			hcl.Locals{"third_local"},
		},
		Data: []hcl.Data{
			hcl.Data{Name: "a_data", Type: "consul_key_prefix"},
		},
		Providers: []hcl.Provider{
			hcl.Provider{Name: "google", Version: "=1.28.0"},
		},
		Terraform: hcl.Terraform{Version: "> 0.12.0", Backend: "gcs"},
		Modules: []hcl.Module{
			hcl.Module{Name: "consul", Version: "0.0.5"},
			hcl.Module{Name: "network", Version: "1.2.3"},
		},
	},
}

func TestVerifyFile(t *testing.T) {
	// test1 well formed ParsedFile
	parsedFileTest := parsedFile
	authorizedBocks := []string{"module", "locals", "provider", "data",
		"variable", "resource", "output", "terraform"}
	expectedOut := ""
	testOut := capturer.CaptureStdout(func() {
		checks.VerifyFile(parsedFileTest, "^[a-z0-9_]*$", authorizedBocks)
	})

	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyFile(ok) mismatch (-want +got):\n%s", diff)
	}

	// test2 with unwanted blocks and misnamed bloc
	parsedFileTest.Blocks.Data = []hcl.Data{
		hcl.Data{Name: "a-data", Type: "consul_key_prefix"},
	}
	authorizedBocks = []string{"module", "locals", "provider", "data",
		"variable", "output", "terraform"}

	expectedOut = `
ERROR: main.tf misformed:
  Unmatching "^[a-z0-9_]*$" pattern blockname(s):
    - a-data (data)
  Unauthorized block(s):
    - resource
`

	testOut = capturer.CaptureStdout(func() {
		checks.VerifyFile(parsedFileTest, "^[a-z0-9_]*$", authorizedBocks)
	})

	if diff := cmp.Diff(expectedOut, testOut); diff != "" {
		t.Errorf("VerifyFile(ko) mismatch (-want +got):\n%s", diff)
	}
}
