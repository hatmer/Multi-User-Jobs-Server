package jobs

import (
	//"os"
	"fmt"
	"io"
	"math/rand"
	"os/exec"
	"strconv"
	"log"
	"sync"
)

type Job struct {
	CmdStruct *exec.Cmd
	StdOut    io.ReadCloser
	StdErr    io.ReadCloser
	Output    []byte
	OutputErr []byte
	Owner     string
}

func getUUID() string {
	// random number for simplicity but should be a UUID
	return strconv.Itoa(rand.Intn(100000))
}



func Start(manager map[string]Job, command string, owner string) (string, string) {
	// TODO split command on spaces
	cmd := exec.Command(command) //, "-l")

	stderrIn, _ := cmd.StderrPipe()
	stdoutIn, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
	    log.Fatalf("cmd.Start() failed with '%s'\n", err)
		return "", err.Error()
	}
	//go cmd.Wait()
	var errStdout, errStderr error
	var stdout_copy, stderr_copy []byte
	streamStdOutR, streamStdOutW := io.Pipe()
	streamStdErrR, streamStdErrW := io.Pipe()

    var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		stdout_copy, errStdout = copyAndCapture(streamStdOutW, stdoutIn)
		//go cmd.Wait()
		//wg.Done()
		//wg.Wait()
	}()
    wg.Add(1)
    go func() {
	    stderr_copy, errStderr = copyAndCapture(streamStdErrW, stderrIn)
    }()
	

	//err = cmd.Wait()



//Output: output, ErrOutput: errOutput

        data := Job{CmdStruct: cmd, StdOut: streamStdOutR, StdErr: streamStdErrR, Output: stdout_copy, OutputErr: stderr_copy, Owner: owner}

	// generate an ID and make sure it is unique
	id := getUUID()
	id = "1" // TODO fix
	//for manager[id] != nil {
	//	id = getUUID()
	//}

	manager[id] = data
	return id, ""

}

// https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
func copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
    var out []byte
    buf := make([]byte, 1024, 1024)
    for {
        n, err := r.Read(buf[:])   // read from reader and store in buffer
        if n > 0 {
            d := buf[:n]           
            out = append(out, d...)  // copy everything to out
            _, err := w.Write(d)     // and then write it to w
            if err != nil {
                return out, err
            }
        }
        if err != nil {
            // Read returns io.EOF at the end of file, which is not an error for us
            if err == io.EOF {
                err = nil
            }
            return out, err
        }
    }
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
