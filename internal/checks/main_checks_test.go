package checks_test

import (
	"flag"
	"github.com/google/go-cmp/cmp"
	"github.com/kami-zh/go-capturer"
	"github.com/thazelart/terraform-validator/internal/checks"
	"github.com/thazelart/terraform-validator/internal/config"
	"io/ioutil"
	"os"
	"testing"
)

type tc struct {
	Path   string
	Golden string
}

var (
	update    = flag.Bool("update", false, "update .golden.json files")
	testCases = []tc{
		{
			Path:   "testdata/ko_custom_config",
			Golden: "testdata/ko_custom_config.golden",
		},
		{
			Path:   "testdata/ko_default_config",
			Golden: "testdata/ko_default_config.golden",
		},
		{
			Path:   "testdata/ok_custom_config",
			Golden: "testdata/ok_custom_config.golden",
		},
		{
			Path:   "testdata/ok_default_config",
			Golden: "testdata/ok_default_config.golden",
		},
		{
			Path:   "testdata/recurse_cust_default",
			Golden: "testdata/recurse_cust_default.golden",
		},
		{
			Path:   "testdata/recurse_cust_empty_nothing",
			Golden: "testdata/recurse_cust_empty_nothing.golden",
		},
		{
			Path:   "testdata/recurse_cust_nothing",
			Golden: "testdata/recurse_cust_nothing.golden",
		},
		{
			Path:   "testdata/recurse_default_cust",
			Golden: "testdata/recurse_default_cust.golden",
		},
		{
			Path:   "testdata/recurse_default_emptywithconf_cust",
			Golden: "testdata/recurse_default_emptywithconf_cust.golden",
		},
		{
			Path:   "testdata/recurse_override_default",
			Golden: "testdata/recurse_override_default.golden",
		},
	}
)

func TestMaintChecks(t *testing.T) {
	for _, testCase := range testCases {
		os.Args = []string{"terraform-validator", testCase.Path}
		rootDir := config.ParseArgs("dev")

		response := capturer.CaptureStdout(func() {
			checks.MainChecks(config.DefaultTfvConfig(), rootDir)
		})

		goldenFile := testCase.Golden
		if *update {
			ioutil.WriteFile(goldenFile, []byte(response), 0644)
		}
		expected, _ := ioutil.ReadFile(goldenFile)

		if diff := cmp.Diff(string(expected), response); diff != "" {
			t.Errorf("MaintChecks(%s) mismatch (-want +got):\n%s", testCase.Path, diff)
		}
	}
}
