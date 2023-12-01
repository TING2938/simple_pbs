package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"simple_pbs/util"
	"strings"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func print_help() {
	fmt.Printf("help: qstat jobs.pbs ...")
}

func main() {
	nargs := len(os.Args)
	fnm := ""
	args := ""
	switch {
	case nargs < 2:
		print_help()
		return
	case nargs == 2:
		fnm = os.Args[1]
	case nargs > 2:
		fnm = os.Args[1]
		args = strings.Join(os.Args[2:], " ")
	}

	metadata := util.PBS_metadata{}
	metadata.Parse(fnm)

	cgPath := util.Get_cgroup_path(metadata.Task_name)
	control, err := cgroups.New(util.Subsystem([]cgroups.Name{cgroups.Cpu}), cgroups.StaticPath(cgPath), &specs.LinuxResources{})
	if err != nil {
		log.Fatal(err)
		return
	}

	command := exec.Command("bash", "-c", fmt.Sprintf(". %s %s", fnm, args))

	command.Start()
	if err = control.AddTask(cgroups.Process{Pid: command.Process.Pid}); err != nil {
		log.Fatal(err)
		return
	}

	tasks, err := control.Tasks(cgroups.Cpu, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current tasks: %v", tasks)
}
