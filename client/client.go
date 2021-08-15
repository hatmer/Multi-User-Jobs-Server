/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	pb "project/proto"
	"strings"
	"time"
)

const serverAddr = "127.0.0.1:50051"

// stream streams output of a job
func stream(client pb.JobClient, req *pb.JobControlRequest) {
	log.Printf("streaming")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // TODO timeout for stream?
	defer cancel()
	stream, err := client.Stream(ctx, req)
	if err != nil {
		log.Fatalf("%v stream fxn error, %v", client, err)
	}
	for {
		line, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v stream recv error, %v", client, err)
		}
		/*lines := strings.Split(line.GetText(), "\n")
		for i := 0; i < len(lines); i++ {
			if len(lines[i]) > 0 {
				fmt.Println(lines[i])
			}
		}*/
		printOutput(line.GetText())
	}
	log.Printf("stream complete")
}

func printOutput(s string) {
    lines := strings.Split(s, "\n")
		for i := 0; i < len(lines); i++ {
			if len(lines[i]) > 0 {
				fmt.Println(lines[i])
			}
		}
}

// starts a job
func start(client pb.JobClient, req *pb.JobStartRequest) {
	//log.Printf("Looking for features within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Start(ctx, req)
	if err != nil {
		log.Fatalf("%v start fxn err, %v", client, err)
	}
	log.Println(resp.GetStatus())
}

// stops a job
func stop(client pb.JobClient, req *pb.JobControlRequest) {
	log.Printf("Stopping job")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Stop(ctx, req)
	if err != nil {
		log.Fatalf("%v stop fxn err, %v", client, err)
	}
	log.Println(resp.GetStatus())
}

// gets status of a job
func status(client pb.JobClient, req *pb.JobControlRequest) {
	log.Println("Status")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Status(ctx, req)
	if err != nil {
		log.Fatalf("%v stop fxn err, %v", client, err)
	}
	log.Println(resp.GetStatus())
}

// gets output of a completed job
func output(client pb.JobClient, req *pb.JobControlRequest) {
	log.Printf("Requesting job output")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.Output(ctx, req)
	if err != nil {
		log.Fatalf("%v output fxn err, %v", client, err)
	}
	printOutput(resp.GetStatus())
	//log.Println(resp.GetStatus())
}

func main() {

	// TODO read args: start/stop/status/stream jobID

	// Load the client certificate and its key
	clientCert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		log.Fatalf("Failed to load client certificate and key. %s.", err)
	}

	// Load the CA certificate
	trustedCert, err := ioutil.ReadFile("cacert.pem")
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
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create a new TLS credentials based on the TLS configuration
	cred := credentials.NewTLS(tlsConfig)

	// Dial the gRPC server with the given credentials
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(cred))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			log.Printf("Unable to close gRPC channel. %s.", err)
		}
	}()

	client := pb.NewJobClient(conn)

	// Looking for a valid feature
	start(client, &pb.JobStartRequest{Job: "./test.sh"})

	// Looking for features between 40, -75 and 42, -73.
	status(client, &pb.JobControlRequest{JobID: "1", Request: "status"})
	stream(client, &pb.JobControlRequest{JobID: "1", Request: "stream"})

	status(client, &pb.JobControlRequest{JobID: "1", Request: "status"})
	//       stop(client, &pb.JobControlRequest{JobID: "1", Request: "stop"})
	//	status(client, &pb.JobControlRequest{JobID: "1", Request: "status"})
}
