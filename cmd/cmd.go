package cmd

import (
	"fmt"
	"os/exec"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var cmdCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Interactive docker command selector",
	Run: func(cmd *cobra.Command, args []string) {
		mainMenu()
	},
}

func init() {
	rootCmd.AddCommand(cmdCmd)
}

func mainMenu() {
	options := []string{"ps", "exec", "rm", "rmi", "stop", "start"}
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

func runDockerRm() {
	fmt.Println("rm: 実装予定です")
}

func runDockerRmi() {
	fmt.Println("rmi: 実装予定です")
}

func runDockerStop() {
	fmt.Println("stop: 実装予定です")
}

func runDockerStart() {
	fmt.Println("start: 実装予定です")
}
