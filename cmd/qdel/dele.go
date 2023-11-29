package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/containerd/cgroups"
)

func main() {
	var cgPath = flag.String("cgroup_path", "", "cg-path is cgroup path name")

	flag.Parse()

	for {
		control, err := cgroups.Load(cgroups.V1, cgroups.StaticPath(*cgPath))
		if err != nil {
			log.Printf("load cgroups: %v", err)
			return
		}
		tasks, err := control.Tasks(cgroups.Freezer, false)
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
