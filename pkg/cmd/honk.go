package cmd

import (
	"bytes"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"

	"github.com/honk-ci/kubectl-honk/pkg/utils"
	"github.com/qeesung/image2ascii/convert"
	"github.com/spf13/cobra"

	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd/api"
)

var (
	honkExample = `
	# honk!
	%[1]s honk

	# Add your custom Goose Image
	%[1]s honk --goose <PATH_GOOSE_IMAGE>
`
)

// HonkOptions provides information required to update
// the current context on a user's KUBECONFIG
type HonkOptions struct {
	configFlags *genericclioptions.ConfigFlags

	resultingContext *api.Context

	rawConfig     api.Config
	gooseFilePath string
	args          []string

	genericclioptions.IOStreams
}

// NewHonkOptions provides an instance of NamespaceOptions with default values
func NewHonkOptions(streams genericclioptions.IOStreams) *HonkOptions {
	return &HonkOptions{
		configFlags: genericclioptions.NewConfigFlags(true),

		IOStreams: streams,
	}
}

// NewCmdHonk provides a cobra command wrapping NewHonkOptions
func NewCmdHonk(streams genericclioptions.IOStreams) *cobra.Command {
	o := NewHonkOptions(streams)

	cmd := &cobra.Command{
		Use:          "honk [flags]",
		Short:        "Honk the Planet",
		Example:      fmt.Sprintf(honkExample, "kubectl"),
		SilenceUsage: true,
		RunE: func(c *cobra.Command, args []string) error {
			if err := o.Run(); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&o.gooseFilePath, "goose", "", "path for a custom goose image")

	return cmd
}

// Run the honk command honk!
func (o *HonkOptions) Run() error {

	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.Ratio = 0.2
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()

	var gooseASCII string
	if o.gooseFilePath != "" {
		gooseASCII = converter.ImageFile2ASCIIString(o.gooseFilePath, &convertOptions)
	} else {
		goose, _, err := image.Decode(bytes.NewReader(utils.GetDefaultGoose()))
		if err != nil {
			return err
		}
		gooseASCII = converter.Image2ASCIIString(goose, &convertOptions)
	}

	fmt.Fprintf(o.Out, "Honk Honk!\n%s\n", gooseASCII)

	return nil
}
