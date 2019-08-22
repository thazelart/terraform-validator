package tfv_test

import (
	"flag"
	"github.com/google/go-cmp/cmp"
	"github.com/kami-zh/go-capturer"
	"github.com/thazelart/terraform-validator/internal/config"
	"github.com/thazelart/terraform-validator/internal/tfv"
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
		tc{
			Path:   "../../testdata/ko_custom_config",
			Golden: "../../testdata/ko_custom_config.golden",
		},
		tc{
			Path:   "../../testdata/ko_default_config",
			Golden: "../../testdata/ko_default_config.golden",
		},
		tc{
			Path:   "../../testdata/ok_custom_config",
			Golden: "../../testdata/ok_custom_config.golden",
		},
		tc{
			Path:   "../../testdata/ok_default_config",
			Golden: "../../testdata/ok_default_config.golden",
		},
	}
)

func TestMaintChecks(t *testing.T) {
	for _, testCase := range testCases {
		os.Args = []string{"terraform-validator", testCase.Path}
		globalConfig := config.GenerateGlobalConfig("dev")

		response := capturer.CaptureStdout(func() {
			tfv.MainChecks(globalConfig)
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
