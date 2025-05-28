package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerRm()
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
}

func runDockerRm() {
	cmdContainers := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}\t{{.Image}}\t{{.Names}}\t{{.Status}}\t{{.CreatedAt}}")
	out, err := cmdContainers.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get docker containers: %v\n", err)
		return
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
		fmt.Println("No docker containers found.")
		return
	}
	// map: 表示行 -> ID
	lineToID := make(map[string]string)
	for _, line := range lines {
		fields := strings.SplitN(line, "\t", 2)
		if len(fields) < 2 {
			continue
		}
		id := fields[0]
		lineToID[line] = id
	}

	container, err := internal.SelectWithPeco(lines)
	if err != nil {
		fmt.Println("peco container select error:", err)
		return
	}

	cmdRm := exec.Command("docker", "rm", "-f", lineToID[container])
	cmdRm.Stdin = os.Stdin
	cmdRm.Stdout = os.Stdout
	cmdRm.Stderr = os.Stderr
	fmt.Println("Removing container:", lineToID[container])
	cmdRm.Run()
}
