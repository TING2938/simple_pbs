package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

func main() {
	var task_name = flag.String("task", "", "task name")
	var period = flag.Uint64("cpu_period", 100000, "cpu limit value, default is 100% ")
	var quota = flag.Int64("cpu_quota", -1, "cpu limit value, default is 100% ")
	var cmd = flag.String("cmd", "", "your application cmd")

	flag.Parse()

	cgPath := fmt.Sprintf("pbs_%v", task_name)
	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath(cgPath), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota:  quota,
			Period: period,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	command := run(*cmd)
	pid := command.Process.Pid
	if err = control.AddTask(cgroups.Process{Pid: pid}); err != nil {
		log.Fatal(err)
		return
	}

	tasks, err := control.Tasks(cgroups.Freezer, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current tasks: %v", tasks)
}
