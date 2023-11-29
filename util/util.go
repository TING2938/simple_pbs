package util

import (
	"fmt"
	"slices"

	"github.com/containerd/cgroups"
)

const (
	KB = 1024
	MB = 1024 * KB
)

func Subsystem(names []cgroups.Name) cgroups.Hierarchy {
	return func() ([]cgroups.Subsystem, error) {
		subsystems, err := cgroups.V1()
		if err != nil {
			return nil, err
		}
		var enabled []cgroups.Subsystem
		for _, s := range subsystems {
			if slices.Contains(names, s.Name()) {
				enabled = append(enabled, s)
			}
		}
		return enabled, nil
	}
}

func Get_cgroup_path(task_name string) string {
	return fmt.Sprintf("pbs_%v", task_name)
}
