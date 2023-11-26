package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/containerd/cgroups"
	"github.com/opencontainers/runtime-spec/specs-go"
)

const (
	kb = 1024
	mb = 1024 * kb
)

func main() {
	var cgPath = flag.String("cgroup_path", "", "cg-path is cgroup path name")
	var period = flag.Uint64("cpu_period", 100000, "cpu limit value, default is 100% ")
	var quota = flag.Int64("cpu_quota", -1, "cpu limit value, default is 100% ")
	var memLimit = flag.Int("mem_limit", 100, "mem limit value, default is 100mb ")

	var cmd = flag.String("cmd", "", "your application cmd")
	var args = flag.String("args", "", "cmd args")

	flag.Parse()

	cpuLimit := float32(*quota) / float32(*period) * 100
	limit := int64(*memLimit * mb)

	log.Printf("cgroup_path: %s, cpu_quota: %v, cpu_period: %v,max (%v%%), mem_limit: %vm (%d), cmd: %s, args: %v \n",
		*cgPath, *quota, *period, cpuLimit, *memLimit, limit, *cmd, *args)

	control, err := cgroups.New(cgroups.V1, cgroups.StaticPath(*cgPath), &specs.LinuxResources{
		CPU: &specs.LinuxCPU{
			Quota:  quota,
			Period: period,
		},
		Memory: &specs.LinuxMemory{
			Limit: &limit,
		},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	defer control.Delete()

	pid, command := run(*cmd, strings.Split(*args, " ")...)
	log.Printf("run process done, pid: %v, add to cgroup task\n", pid)
	if err = control.AddTask(cgroups.Process{Pid: pid}); err != nil {
		log.Fatal(err)
		return
	}

	tasks, err := control.Tasks(cgroups.Freezer, false)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Current tasks: %v", tasks)

	command.Wait()
}

func run(cmd string, args ...string) (int, *exec.Cmd) {
	log.Printf("[run], cmd: %s, args: %v", cmd, args)
	command := exec.Command(cmd, args...)

	command.Stdout = os.Stdout
	command.Stderr = os.Stdout

	err := command.Start()
	if err != nil {
		log.Fatalf("Start error, %v", err)
		return 0, command
	}
	for {
		if command.Process != nil {
			return command.Process.Pid, command
		}
		time.Sleep(1000 * time.Microsecond)
	}
}
