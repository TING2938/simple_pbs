package main

import (
	"flag"
	"log"
	"os/exec"
	"simple_pbs/util"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {

	var task_name = flag.String("task", "", "task name")
	var cmd = flag.String("cmd", "", "your application cmd")

	flag.Parse()

	cgPath := util.Get_cgroup_path(*task_name)
	control, err := cgroups.New(util.Subsystem([]cgroups.Name{cgroups.Cpu}), cgroups.StaticPath(cgPath), &specs.LinuxResources{})
	if err != nil {
		log.Fatal(err)
		return
	}

	command := exec.Command("bash", "-c", *cmd)
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
