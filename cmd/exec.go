package cmd

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in a running Docker container",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerExec()
	},
}

func runDockerExec() {
	// Step 1: List running containers
	psCmd := exec.Command("docker", "ps", "--format", "{{.ID}}\t{{.Names}}\t{{.Image}}\t{{.Status}}")
	psOut, err := psCmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to list containers: %v\n", err)
		os.Exit(1)
	}
	if len(psOut) == 0 {
		fmt.Println("No running containers found.")
		return
	}

	// Step 2: Select container via peco
	pecoCmd := exec.Command("peco")
	pecoCmd.Stdin = bytes.NewReader(psOut)
	var pecoOut bytes.Buffer
	pecoCmd.Stdout = &pecoOut
	if err := pecoCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run peco: %v\n", err)
		os.Exit(1)
	}

	selectedLine := strings.TrimSpace(pecoOut.String())
	if selectedLine == "" {
		fmt.Println("No container selected.")
		return
	}
	fields := strings.Split(selectedLine, "\t")
	if len(fields) < 1 {
		fmt.Println("Invalid selection.")
		return
	}
	containerID := fields[0]

	// Step 3: Load commands from config
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get home directory: %v\n", err)
		os.Exit(1)
	}
	configPath := filepath.Join(homeDir, ".dok.config")
	commands, err := internal.LoadCommands(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load commands: %v\n", err)
		os.Exit(1)
	}
	if len(commands) == 0 {
		fmt.Println("No commands found in config.")
		return
	}

	// Prepare commands list for selection
	commandsStr := strings.Join(commands, "\n")

	// Step 4: Select command via peco
	pecoCmd = exec.Command("peco")
	pecoCmd.Stdin = strings.NewReader(commandsStr)
	pecoOut.Reset()
	pecoCmd.Stdout = &pecoOut
	if err := pecoCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run peco for commands: %v\n", err)
		os.Exit(1)
	}
	selectedCommand := strings.TrimSpace(pecoOut.String())
	if selectedCommand == "" {
		fmt.Println("No command selected.")
		return
	}
	commandToRun := selectedCommand

	// Step 5: Select interactive flag
	flags := []string{"-i", "-ti", "-t", "none"}
	flagsStr := strings.Join(flags, "\n")
	pecoCmd = exec.Command("peco")
	pecoCmd.Stdin = strings.NewReader(flagsStr)
	pecoOut.Reset()
	pecoCmd.Stdout = &pecoOut
	if err := pecoCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run peco for flags: %v\n", err)
		os.Exit(1)
	}
	selectedFlag := strings.TrimSpace(pecoOut.String())
	if selectedFlag == "" {
		fmt.Println("No flag selected.")
		return
	}

	// Step 6: Run docker exec
	argsExec := []string{"exec"}
	if selectedFlag != "none" {
		argsExec = append(argsExec, selectedFlag)
	}
	argsExec = append(argsExec, containerID)
	argsExec = append(argsExec, strings.Split(commandToRun, " ")...)

	dockerExecCmd := exec.Command("docker", argsExec...)
	dockerExecCmd.Stdin = os.Stdin
	dockerExecCmd.Stdout = os.Stdout
	dockerExecCmd.Stderr = os.Stderr
	if err := dockerExecCmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run docker exec: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(execCmd)
}
