package util

import (
	"log"
	"os/exec"
	"time"
)

const (
	kb = 1024
	mb = 1024 * kb
)

func Run(command string) *exec.Cmd {
	log.Printf("[run], cmd: %s", command)
	cmd := exec.Command("bash", "-c", command)

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Start error, %v", err)
		return cmd
	}
	for {
		if cmd.Process != nil {
			return cmd
		}
		time.Sleep(1000 * time.Microsecond)
	}
}
