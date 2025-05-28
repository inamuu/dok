package cmd

import (
	"fmt"
	"os/exec"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Interactive docker command selector",
	Run: func(cmd *cobra.Command, args []string) {
		mainMenu()
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)
}

func mainMenu() {
	options := []string{"ps", "exec", "rm", "rmi", "run", "stop", "start"}
	choice, err := internal.SelectWithPeco(options)
	if err != nil {
		fmt.Println("peco error:", err)
		return
	}

	switch choice {
	case "ps":
		runDockerPs()
	case "exec":
		runDockerExec()
	case "rm":
		runDockerRm()
	case "rmi":
		runDockerRmi()
	case "run":
		runDockerRun()
	case "stop":
		runDockerStop()
	case "start":
		runDockerStart()
	}
}

func runDockerPs() {
	cmd := exec.Command("docker", "ps")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = internal.StdStreams()
	cmd.Run()
}

func runDockerExec() {
	fmt.Println("exec: 実装予定です")
}

func runDockerStop() {
	fmt.Println("stop: 実装予定です")
}

func runDockerStart() {
	fmt.Println("start: 実装予定です")
}
