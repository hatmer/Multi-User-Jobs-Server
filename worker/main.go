package worker

import (
    "os/exec"
    "os"
)
/*
func Start(command string) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
	cmd := exec.Command(command)
	stdout, err := cmd.StdoutPipe()

        // while the process continues, read from stdout and stderr into channel

        for _, n := range numbers {
            out <- n
        }
    }()
    return out
}
*/
func Run(command string) string {
	cmd := exec.Command(command) //, "-l")
	// TODO check for error
	cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
        cmd.Run()
	return "ok"
}
