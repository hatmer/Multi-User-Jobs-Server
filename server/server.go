package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"io/ioutil"
	"log"
	"net"
	"project/jobs"
	pb "project/proto"
)

const (
	port     = ":50051"
	certFile = "creds/server.crt"
	keyFile  = "creds/server.pem"
)

type server struct {
	pb.UnimplementedJobServer
	manager map[string]jobs.Job
}

// Start a job
func (s *server) Start(ctx context.Context, in *pb.JobStartRequest) (*pb.JobInfo, error) {
	// TODO input sanitization
	job := in.GetJob()
	log.Printf("Received start request: %v", job)
	//p, ok := peer.FromContext(ctx)
	//log.Printf("peer info: %v, %v", p, ok)
	jobID, err := jobs.Start(s.manager, job, "owner")
	//log.Printf("JobID, Result: %v, %v", jobID, res)
	res := "job started"
	if err != nil {
	    res = "job failed to start"
	}
	return &pb.JobInfo{JobID: jobID, Response: res}, err
}

// Stop a job
func (s *server) Stop(ctx context.Context, in *pb.JobControlRequest) (*pb.JobInfo, error) {
    // TODO input sanitization
    jobID := in.GetJobID()
	log.Printf("Received stop request for job: %v", jobID)
	//p, ok := peer.FromContext(ctx)
	// TODO verify ownership
	//log.Printf("peer info: %v, %v", p, ok)
	//owner := "owner"
	res, err := jobs.Stop(s.manager, jobID)
	log.Printf("Job stop result, %v, %v", jobID, res)

	return &pb.JobInfo{JobID: jobID, Response: res}, err // TODO better error passing around
}

// Get status of a job
func (s *server) Status(ctx context.Context, in *pb.JobControlRequest) (*pb.JobInfo, error) {
    // TODO input sanitization
	jobID := in.GetJobID()
	log.Printf("Received status request for job: %s", jobID)
	//p, ok := peer.FromContext(ctx)
	// TODO verify ownership
	//log.Printf("peer info: %v, %v", p, ok)

	res, err := jobs.Status(s.manager, jobID)
	log.Printf("Job status result, %v, %v", jobID, res)

	return &pb.JobInfo{JobID: jobID, Response: res}, err
}

// Get final output of a job
func (s *server) Output(ctx context.Context, in *pb.JobControlRequest) (*pb.JobInfo, error) {
    // TODO input sanitization
	jobID := in.GetJobID()
	log.Printf("Received output request for job: %s", jobID)
	//p, ok := peer.FromContext(ctx)
	// TODO verify ownership
	//log.Printf("peer info: %v, %v", p, ok)
err := nil
	res := "no output yet: job is still running"
	job, exists := s.manager[jobID]
	if !exists {
		res = "job does not exist"
		err = error.New("invalid job ID")
	} else {

		if job.CmdStruct.ProcessState != nil {
			res = string(*job.Output)
			res = res + "\n" + string(job.CmdStruct.ProcessState.ExitCode())
		}
	}
	//log.Printf("Job output result, %v, %v", jobID, res)

	return &pb.JobInfo{JobID: jobID, Response: res}, err

}

// stream output of a job
func (s *server) Stream(in *pb.JobControlRequest, stream pb.Job_StreamServer) error {
    // TODO input sanitization
    // TODO verify ownership
	jobID := in.GetJobID()
	log.Printf("Received stream request for job: %s", jobID)
	job := s.manager[jobID]
	output := make([]byte, 1024)

	for job.CmdStruct.ProcessState == nil { // while the process is still running
		log.Println("streaming...")
		n, _ := job.StdOut.Read(output) // TODO handle stderr also
		log.Printf("read %d bytes of output", n)
		if n > 0 {
			if err := stream.Send(&pb.Line{Text: string(output)[:n]}); err != nil {
				return err
			}
		}
	}
	log.Println("process completed")
	n, _ := job.StdOut.Read(output) // TODO handle stderr also
	log.Printf("read final %d bytes of output", n)
	if n > 0 {
		if err := stream.Send(&pb.Line{Text: string(output)}); err != nil {
			return err
		}
	}
	ret := fmt.Sprintf("Job exited with code: %d", job.CmdStruct.ProcessState.ExitCode())
	if err := stream.Send(&pb.Line{Text: ret}); err != nil {
		return err
	}

	return nil
}

func main() {
	// Load the server certificate and its key
	serverCert, err := tls.LoadX509KeyPair("creds/server.pem", "creds/server.key")
	if err != nil {
		log.Fatalf("Failed to load server certificate and key. %s.", err)
	}

	// Load the CA certificate
	trustedCert, err := ioutil.ReadFile("creds/cacert.pem")
	if err != nil {
		log.Fatalf("Failed to load trusted certificate. %s.", err)
	}

	// Put the CA certificate to certificate pool
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		log.Fatalf("Failed to append trusted certificate to certificate pool. %s.", err)
	}

	// Create the TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)

	// Create a listener that listens to localhost
	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to start listener. %s.", err)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			log.Printf("Failed to close listener. %s\n", err)
		}
	}()

	// Create a new gRPC server
	s := grpc.NewServer(grpc.Creds(cred))
	pb.RegisterJobServer(s, &server{manager: make(map[string]jobs.Job)})

	// Start the gRPC server
	log.Printf("server listening at localhost:%v", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatalf("Failed to start gRPC server. %s.", err)
	}
}
