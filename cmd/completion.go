package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completionTarget string

// completionCmd represents the completion command.
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completion script for " + rootCmd.Use,
	Long: `Generates a shell completion script for ` + rootCmd.Use + `.
	NOTE: The current version supports Bash only.
		  This should work for *nix systems with Bash installed.

	By default, the file is written directly to /etc/bash_completion.d
	for convenience, and the command may need superuser rights, e.g.:

		sudo ` + rootCmd.Use + ` completion

	Add ` + "`--completionfile=/path/to/file`" + ` flag to set alternative
	file-path and name.

	For e.g. on OSX with bash completion installed with brew you should

	` + rootCmd.Use + ` completion --completionfile /etc/bash_completion.d/` + rootCmd.Use + `.sh

	or using if using brew

	` + rootCmd.Use + ` completion --completionfile $(brew --prefix)/etc/bash_completion.d/` + rootCmd.Use + `.sh

	Logout and in again to reload the completion scripts,
	or just source them directly:

		. /etc/bash_completion

	or using if using brew

		. $(brew --prefix)/etc/bash_completion`,

	Run: Completion,
}

// Completion is a helper function to allow passing arguments to
// other functions (so that they can be unit tested).
func Completion(cmd *cobra.Command, args []string) {
	err := cmd.Root().GenBashCompletionFile(completionTarget)
	completion(err)
}

func completion(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bash completion file for "+rootCmd.Use+" saved to", completionTarget)
}

func init() { //nolint:gochecknoinits
	rootCmd.AddCommand(completionCmd)

	completionCmd.PersistentFlags().StringVarP(&completionTarget,
		"completionfile",
		"",
		"/etc/bash_completion.d/"+rootCmd.Use+".sh",
		"completion file")
	// Required for bash-completion
	_ = completionCmd.PersistentFlags().SetAnnotation("completionfile", cobra.BashCompFilenameExt, []string{})
}
