package jobs

import (
	"os"
	"os/exec"
    "strconv"
	"math/rand"
)

func getUUID() string {
	return strconv.Itoa(rand.Intn(100000))
}

func Start(manager map[string](*exec.Cmd), command string) (string, string) {
    // TODO split command on spaces
	cmd := exec.Command(command) //, "-l")
	
	// TODO use pipes (are they buffered? maybe use channels?)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
	    return "", err.Error()
	}

    // generate an ID and make sure it is unique
	id := getUUID()
	//for manager[id] != nil {
	//	id = getUUID()
	//}
	
	manager[id] = cmd
	return id, ""
	
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
