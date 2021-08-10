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

// Package main implements a server for API service.
package main

import (
	"context"
	"log"
	"net"
	"project/worker"

	"google.golang.org/grpc"
	pb "project/proto"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedJobServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) Start(ctx context.Context, in *pb.JobStartRequest) (*pb.JobStatus, error) {
	log.Printf("Received: %v", in.GetJob())
	// TODO input sanitization?
	res := worker.Run(in.getJob())
	
	return &pb.JobStatus{jobID: "1", status: "probably ok"}, nil
}

/*
func (s *routeGuideServer) StreamOutput(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
  for _, feature := range s.savedFeatures {
    if inRange(feature.Location, rect) {
      if err := stream.Send(feature); err != nil {
        return err
      }
    }
  }
  return nil
}
*/


func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterJobServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
