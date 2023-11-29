package util

import "fmt"

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
