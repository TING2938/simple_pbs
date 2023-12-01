package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Worker struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}

type Nodes struct {
	Nodes map[string]Worker `json:"nodes"`
}

func (w *Worker) Command() string {
	return fmt.Sprintf("ssh -p %d root@%s", w.Port, w.IP)
}

func (n *Nodes) Worker(name string) Worker {
	return n.Nodes[name]
}

type PBS_metadata struct {
	Nodes         []string
	GPU_per_nodes int
	Master_port   int
	Num_nodes     int
	Task_name     string
	Work_dir      string
}

func (m *PBS_metadata) Parse(fnm string) error {
	content, err := os.ReadFile(fnm)
	if err != nil {
		return err
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "#PBS") {
			all_split := strings.Split(line, " ")[1:]
			switch all_split[0] {
			case "--NODES":
				m.Nodes = all_split[1:]
			case "--GPU_PER_NODES":
				m.GPU_per_nodes, err = strconv.Atoi(all_split[1])
			case "--MASTER_PORT":
				m.Master_port, err = strconv.Atoi(all_split[1])
			case "--NUM_NODES":
				m.Num_nodes, err = strconv.Atoi(all_split[1])
			case "--NAME":
				m.Task_name = all_split[1]
			case "--WORKDIR":
				m.Work_dir = all_split[1]
				if strings.ToLower(m.Work_dir) == "pwd" {
					m.Work_dir, err = os.Getwd()
				}
			}
		}
	}
	return err
}
