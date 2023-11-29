package func_test

import (
	"fmt"
	"os"
	"simple_pbs/util"
	"strings"
	"testing"
)

func Test_load_config(t *testing.T) {
	fnm := "../tests/pbs_config.json"
	outputData := util.Nodes{}
	util.Load_json(fnm, &outputData)
	fmt.Print(outputData)
}

func Test_parse_metadata(t *testing.T) {
	fnm := "../tests/jobs_template.pbs"
	content, err := os.ReadFile(fnm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	lines := strings.Split(string(content), "\n")
	fmt.Printf("content: %v\n", lines)
}
