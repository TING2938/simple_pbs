package func_test

import (
	"fmt"
	"simple_pbs/util"
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
	meta := util.PBS_metadata{}
	meta.Parse(fnm)
	fmt.Printf("meta: %v\n", meta)
}
