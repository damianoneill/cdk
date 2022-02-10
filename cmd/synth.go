package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

var stack = "cdk"

var synthCommand = &cobra.Command{
	Use:   "synth",
	Short: "Output the current build information",
	Long:  "Version, Commit and Date will be output from the Build Info.",
	Run: func(cmd *cobra.Command, args []string) {
		Synth(os.Stdout)
	},
}

func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// docker.NewDockerProvider(stack, jsii.String("provider"), &docker.DockerProviderConfig{})

	// dockerImage := docker.NewImage(stack, jsii.String("nginxImage"), &docker.ImageConfig{
	// 	Name:        jsii.String("nginx:latest"),
	// 	KeepLocally: jsii.Bool(false),
	// })

	// docker.NewContainer(stack, jsii.String("nginxContainer"), &docker.ContainerConfig{
	// 	Image: dockerImage.Latest(),
	// 	Name:  jsii.String("tutorial"),
	// 	Ports: &[]*docker.ContainerPorts{{
	// 		Internal: jsii.Number(80), External: jsii.Number(8000),
	// 	}},
	// })

	return stack
}

// Version outputs a formatted version message to the passed writer.
func Synth(w io.Writer) {
	app := cdktf.NewApp(nil)

	NewMyStack(app, stack)

	app.Synth()
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(synthCommand)
}
