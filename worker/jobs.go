package jobs

import (
    "os/exec"
    "os"
)

type Manager struct {
    jobIDs   []string
}


func (m Manager) Start(command string) (string, string) {
	cmd := exec.Command(command) //, "-l")
	// TODO check for error
	cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    err := cmd.Start()
    
    id := 1 // TODO get a new unique ID
	return id, err
}


func (m Manager) Status(command string) string {
    return "ok"
}

func (m Manager) Stop(command string) string {
    return "ok"
}

/*
func Stream(command string) <-chan int {
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