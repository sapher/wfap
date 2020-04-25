package wfap

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

/**
 * Execute a command
 * Take program name and list of commands to execute
 * Return Stdout result
 */
func executeCmd(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	log.Printf("exec: %s %s\n", name, strings.Join(args, " "))
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	output := strings.TrimSuffix(stdout.String(), "\n")
	return output, nil
}
