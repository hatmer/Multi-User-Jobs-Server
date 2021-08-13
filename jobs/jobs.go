package jobs

import (
	//"os"
	"os/exec"
    "strconv"
	"math/rand"
	"io"
)

type CmdData struct {
	CmdStruct *exec.Cmd
	StdOut io.ReadCloser
	StdErr io.ReadCloser
	Owner string
}

func getUUID() string {
	return strconv.Itoa(rand.Intn(100000))
}

func Start(manager map[string](CmdData), command string, owner string) (string, string) {
    // TODO split command on spaces
	cmd := exec.Command(command) //, "-l")
	
	// TODO use pipes (are they buffered? maybe use channels?)
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
	    return "", err.Error()
	}

	data := CmdData{CmdStruct: cmd, StdOut: stdout, StdErr: stderr, Owner: owner}

    // generate an ID and make sure it is unique
	id := getUUID()
	id = "1" // TODO fix
	//for manager[id] != nil {
	//	id = getUUID()
	//}
	
	manager[id] = data
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
