package jobs

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"os/exec"
	"strings"
	"sync"
)

type Job struct {
	CmdStruct *exec.Cmd
	StdOut    *bytes.Buffer
	StdErr    *bytes.Buffer
	Output    *[]byte
	OutputErr *[]byte
	Owner     string
}

func getUUID() string {
	return uuid.New().String()
}

func Start(manager map[string]Job, command string, owner string) (string, error) {
	//command = "unshare -m -n -p " + command
	args := strings.Split(command, "<magic6789>")
	cmd := exec.Command(args[0])
	cmd.Args = args

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	err := cmd.Start()
	if err != nil {
		log.Printf("cmd.Start() failed with '%s'\n", err)
		return "", err
	}

	var errStdout, errStderr error

	stdout_copy := make([]byte, 1024, 1024)
	stderr_copy := make([]byte, 1024, 1024)

	var stdoutbuf bytes.Buffer
	var stderrbuf bytes.Buffer

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		errStdout = copyAndCapture(&stdoutbuf, stdout_copy, stdoutIn)
	}()
	wg.Add(1)
	go func() {
		errStderr = copyAndCapture(&stderrbuf, stderr_copy, stderrIn)
	}()
	wg.Add(1)
	go cmd.Wait()

	data := Job{CmdStruct: cmd, StdOut: &stdoutbuf, StdErr: &stderrbuf, Output: &stdout_copy, OutputErr: &stderr_copy, Owner: owner}

	// generate an ID and make sure it is unique
	id := getUUID()

	/*for manager[id] != nil {
		id = getUUID()
	}*/

	manager[id] = data
	return id, nil

}

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
func copyAndCapture(b *bytes.Buffer, buf []byte, r io.Reader) error {
	var out []byte
	for {
		n, err := r.Read(buf[:]) // read from reader and store in buffer
		if n > 0 {
			d := buf[:n]
			out = append(out, d...) // copy everything to out
			_, err := b.Write(d)    // and then write it to w
			if err != nil {
				return err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return err
		}
	}
}

func Status(manager map[string]Job, jobID string) (string, error) {
	job, exists := manager[jobID]
	if !exists {
		return "job does not exist", errors.New("invalid job ID")
	}
	status := "running"

	if job.CmdStruct.ProcessState != nil {
		status = fmt.Sprintf("exited with code %d", job.CmdStruct.ProcessState.ExitCode()) // TODO ensure that exit code means process actually exited
	}

	return fmt.Sprintf("status: %s", status), nil
}

func Stop(manager map[string]Job, jobID string) (string, error) {
	job, exists := manager[jobID]
	if !exists {
		return "cannot stop job", errors.New("invalid job ID")
	}

	// attempt to stop job
	if job.CmdStruct.ProcessState != nil {
		return "cannot stop job: job is not running", errors.New("job is not running")
	}

	err := job.CmdStruct.Process.Kill()
	// TODO: kill process group to ensure that any child processes also killed

	if err != nil {
		return "error occured while stopping job", err
	}

	return "job stopped successfully", nil
}
