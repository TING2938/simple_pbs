package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"simple_pbs/util"
	"time"

	"github.com/containerd/cgroups"
)

func main() {
	var task_name = flag.String("task", "", "task name")
	flag.Parse()
	cgPath := util.Get_cgroup_path(*task_name)

	for {
		control, err := cgroups.Load(util.Subsystem([]cgroups.Name{cgroups.Cpu}), cgroups.StaticPath(cgPath))
		if err != nil {
			log.Printf("load cgroups: %v", err)
			return
		}
		tasks, err := control.Tasks(cgroups.Cpu, false)
		if err != nil {
			log.Printf("task err: %v", err)
			return
		}
		log.Printf("tasks: %v", tasks)

		for _, task := range tasks {
			cmd := exec.Command("kill", "-9", fmt.Sprint(task.Pid))
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
		}

		time.Sleep(2 * time.Second)

		err = control.Delete()
		if err != nil {
			log.Printf("err: %v", err)
		} else {
			return
		}
	}
}
