package jobs

import (
	"os"
	"os/exec"

	"math/rand"
)

func getUUID() string {
	return string(rand.Intn(100))
}

func Start(manager map[string]string, command string) (string, string) {
	cmd := exec.Command(command) //, "-l")
	// TODO check for error
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()

	id := getUUID()
	for manager[id] != "" {
		id = getUUID()
	}
	manager[id] = "running"
	if err != nil {
		return id, err.Error()
	} else {
		return id, ""
	}
}

/*
func Status(command string) string {
    return "ok"
}

func Stop(command string) string {
    return "ok"
}
*/

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
