package cmd

import (
	"os/exec"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Remove a Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerPs()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func runDockerPs() {
	choice, err := internal.SelectWithPeco([]string{"running", "all"})
	if err != nil {
		return
	}

	args := []string{"ps"}
	if choice == "all" {
		args = append(args, "-a")
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = internal.StdStreams()
	cmd.Run()
}
