/*
Copyright 2019 Cisco Systems

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"context"
	pb "github.com/cisco-sso/grpc-challenge/pkg/generated"
	"google.golang.org/grpc"
	"log"
	"time"
)

type ReverseOpts struct {
	Data    string
	Address string
}

func Reverse(opts *ReverseOpts) {
	conn, err := grpc.Dial(opts.Address, grpc.WithInsecure(), grpc.WithTimeout(15*time.Second))

	if err != nil {
		log.Fatalf("Failed to dial %s: %v", opts.Address, err)
	}

	client := pb.NewStringmanipClient(conn)
	resp, err := client.Reversed(context.Background(), &pb.StringRequest{
		Content: opts.Data,
	})
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %v", err)
	}
	println(resp.Content)
}
