package jobs

import (
	//"os"
	"bytes"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"sync"
)

type Job struct {
	CmdStruct *exec.Cmd
	StdOut    *bytes.Buffer //io.ReadCloser
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

	//	stderrIn, _ := cmd.StderrPipe()
	stdoutIn, _ := cmd.StdoutPipe()

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
		return "", err.Error()
	}
	//go cmd.Wait()
	//	var errStdout, errStderr error
	var errStdout error
	stderr_copy := make([]byte 1024)
	//var stdout_copy, stderr_copy []byte
	stdout_copy := make([]byte, 1073741824) // 1 GB max output
	//streamStdOutR, streamStdOutW := io.Pipe()
	var stdoutbuf bytes.Buffer
	//	streamStdErrR, streamStdErrW := io.Pipe()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		errStdout = copyAndCapture(&stdoutbuf, &stdout_copy, stdoutIn)
		log.Printf("stdout copyandcapture returned: %s", stdout_copy)
		//pipeoutput, _ := io.ReadAll(streamStdOutR)
		//log.Printf("output pipe now contains: %s", string(pipeoutput))
		//go cmd.Wait()
		//wg.Done()
		//wg.Wait()
	}()
	wg.Add(1)
	go cmd.Wait()
	//wg.Add(1)
	//go func() {
	//    stderr_copy, errStderr = copyAndCapture(streamStdErrW, stderrIn)
	//}()

	//err = cmd.Wait()

	//Output: output, ErrOutput: errOutput

	data := Job{CmdStruct: cmd, StdOut: &stdoutbuf /*StdErr: streamStdErrR,*/, Output: &stdout_copy, OutputErr: stderr_copy, Owner: owner}

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
func copyAndCapture(b *bytes.Buffer, c *[]byte, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		log.Println("looping")
		n, err := r.Read(buf[:]) // read from reader and store in buffer
		if n > 0 {
			log.Printf("read this from pipe: %s", string(buf))
			d := buf[:n]
			out = append(out, d...) // copy everything to out
			log.Println("writing")
			_, err := b.Write(d) // and then write it to w
			log.Println("write ok")
			if err != nil {
				log.Println("returning on write error")
				return out, err
			}
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				log.Println("got EOF")
				err = nil
			}
			log.Println("returning on read error")
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
