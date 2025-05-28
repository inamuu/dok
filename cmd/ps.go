package cmd

import (
	"os/exec"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Show docker containers",
	Run: func(cmd *cobra.Command, args []string) {
		choice, err := internal.SelectWithPeco([]string{"running", "all"})
		if err != nil {
			return
		}

		args = []string{"ps"}
		if choice == "all" {
			args = append(args, "-a")
		}

		psCmd := exec.Command("docker", args...)
		psCmd.Stdin, psCmd.Stdout, psCmd.Stderr = internal.StdStreams()
		psCmd.Run()
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
}
