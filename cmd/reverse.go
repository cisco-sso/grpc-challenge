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

package main

import (
	"io"

	"github.com/cisco-sso/grpc-challenge/pkg/client"
	"github.com/spf13/cobra"
)

var (
	reverseLong = `reverse a string on the grpc server and return the result`

	reverseExample = `./grpc-challenge reverse "My String"`
	reverseShort   = `reverse a string`
)

func NewCmdReverse(out io.Writer) *cobra.Command {
	options := &client.ReverseOpts{}

	cmd := &cobra.Command{
		Use:     "reverse \"My String\"",
		Short:   reverseShort,
		Long:    reverseLong,
		Example: reverseExample,
		Run: func(cmd *cobra.Command, args []string) {
			client.Reverse(options)
		},
	}
	cmd.Flags().StringVarP(&options.Data, "data", "d", "", "Data to send to the server")
	cmd.MarkFlagRequired("data")

	cmd.Flags().StringVarP(&options.Address, "server", "s", "localhost:50001", "Address of the server")

	return cmd
}
