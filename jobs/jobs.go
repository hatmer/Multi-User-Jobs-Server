package jobs

import (
	//"os"
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"strconv"
)

type Job struct {
	CmdStruct *exec.Cmd
	StdOut    io.ReadCloser
	StdErr    io.ReadCloser
	Owner     string
}

func getUUID() string {
	// random number for simplicity but should be a UUID
	return strconv.Itoa(rand.Intn(100000))
}

func Start(manager map[string]Job, command string, owner string) (string, string) {
	// TODO split command on spaces
	cmd := exec.Command(command) //, "-l")

	// TODO use pipes (are they buffered? maybe use channels?)
	stderr, err := cmd.StderrPipe()
	stdout, err := cmd.StdoutPipe()

	err = cmd.Start()
	if err != nil {
		return "", err.Error()
	}
	go cmd.Wait()

	data := Job{CmdStruct: cmd, StdOut: stdout, StdErr: stderr, Owner: owner}

	// generate an ID and make sure it is unique
	id := getUUID()
	id = "1" // TODO fix
	//for manager[id] != nil {
	//	id = getUUID()
	//}

	manager[id] = data
	return id, ""

}

func Status(manager map[string]Job, jobID string) (string, error) {

	job := manager[jobID]
	status := "running"

	if job.CmdStruct.ProcessState != nil {
		status = fmt.Sprintf("exited with code %d", job.CmdStruct.ProcessState.ExitCode()) // TODO ensure that exit code means process actually exited
	}

	return fmt.Sprintf("jobID %s status: %s", jobID, status), nil
}

func Stop(manager map[string]Job, jobID string) (string, error) {

	job := manager[jobID]

	// attempt to stop job
	if job.CmdStruct.ProcessState != nil {
		return "job already stopped", nil
	}

	err := job.CmdStruct.Process.Kill()
	// TODO: kill process group to ensure that any child processes also killed

	if err != nil {
		return "error occured while stopping job", err
	}

	return "job stopped", nil
}

/* TODO multiple clients can stream same output?
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
