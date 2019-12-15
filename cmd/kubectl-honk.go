package main

import (
	"os"

	"github.com/honk-ci/kubectl-honk/pkg/cmd"
	"github.com/spf13/pflag"

	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-honk", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdHonk(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
