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
	goflag "flag"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	validResources = `

	* reverse
	`
)

var (
	rootLong = `grpc-challenge performs various client server operations`

	rootShort = `grpc-challenge performs various client server operations`
)

type RootCmd struct {
	cobraCommand *cobra.Command
}

var rootCommand = RootCmd{
	cobraCommand: &cobra.Command{
		Use:   "grpc-challenge",
		Short: rootShort,
		Long:  rootLong,
	},
}

func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := rootCommand.cobraCommand.Execute(); err != nil {
		log.Fatalf("Exit unsuccessfully with err: %v", err)
	}
}

func init() {
	NewCmdRoot(os.Stdout)
}

func NewCmdRoot(out io.Writer) *cobra.Command {

	cmd := rootCommand.cobraCommand

	// create subcommands
	cmd.AddCommand(NewCmdReverse(out))
	// Not expected to complete.  The server may be in a different language
	// cmd.AddCommand(NewCmdServe(out))

	return cmd
}

func main() {
	Execute()
}
