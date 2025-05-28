package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Remove a Docker image",
	Run: func(cmd *cobra.Command, args []string) {
		cmdImages := exec.Command("docker", "images", "--format", "{{.ID}}\t{{.Repository}}\t{{.Tag}}")
		out, err := cmdImages.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get docker images: %v\n", err)
			return
		}
		lines := strings.Split(strings.TrimSpace(string(out)), "\n")
		if len(lines) == 0 || (len(lines) == 1 && lines[0] == "") {
			fmt.Println("No docker images found.")
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

		image, err := internal.SelectWithPeco(lines)
		if err != nil {
			fmt.Println("peco image select error:", err)
			return
		}

		cmdRmi := exec.Command("docker", "rmi", "-f", lineToID[image])
		cmdRmi.Stdin = os.Stdin
		cmdRmi.Stdout = os.Stdout
		cmdRmi.Stderr = os.Stderr
		fmt.Println("Removing image:", lineToID[image])
		cmdRmi.Run()
	},
}

func init() {
	rootCmd.AddCommand(rmiCmd)
}
