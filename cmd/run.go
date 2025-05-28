package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/inamuu/dok/internal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Docker container with selected options",
	Run: func(cmd *cobra.Command, args []string) {
		runDockerRun()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runDockerRun() {
	// イメージ一覧取得
	cmdImages := exec.Command("docker", "images", "--format", "{{.Repository}}:{{.Tag}}")
	out, err := cmdImages.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get docker images: %v\n", err)
		return
	}
	images := strings.Split(strings.TrimSpace(string(out)), "\n")
	if len(images) == 0 || (len(images) == 1 && images[0] == "") {
		fmt.Println("No docker images found.")
		return
	}
	image, err := internal.SelectWithPeco(images)
	if err != nil {
		fmt.Println("peco image select error:", err)
		return
	}

	// フラグ選択
	options := []string{"-i", "-t", "-it", "-d", "-dti", "none"}
	selectedFlag, err := internal.SelectWithPeco(options)
	if err != nil {
		fmt.Println("peco flag select error:", err)
		return
	}
	flag := ""
	if selectedFlag != "none" {
		flag = selectedFlag
	}

	// ポート設定
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter source port (host): ")
	srcPort, _ := reader.ReadString('\n')
	fmt.Print("Enter destination port (container): ")
	dstPort, _ := reader.ReadString('\n')
	srcPort = strings.TrimSpace(srcPort)
	dstPort = strings.TrimSpace(dstPort)

	// コマンド選択
	usr, _ := user.Current()
	configPath := filepath.Join(usr.HomeDir, ".dok.config")
	commands, err := internal.LoadCommands(configPath)
	if err != nil {
		fmt.Println("load command error:", err)
		return
	}
	commands = append([]string{""}, commands...)
	selectedCommand, err := internal.SelectWithPeco(commands)
	if err != nil {
		fmt.Println("peco command select error:", err)
		return
	}

	// docker run 実行
	argsRun := []string{"run"}
	if flag != "" {
		argsRun = append(argsRun, flag)
	}
	if srcPort != "" && dstPort != "" {
		argsRun = append(argsRun, "-p", fmt.Sprintf("%s:%s", srcPort, dstPort))
	}
	argsRun = append(argsRun, image)
	if selectedCommand != "" {
		split := strings.Split(selectedCommand, " ")
		argsRun = append(argsRun, split...)
	}

	cmdRun := exec.Command("docker", argsRun...)
	cmdRun.Stdin = os.Stdin
	cmdRun.Stdout = os.Stdout
	cmdRun.Stderr = os.Stderr
	fmt.Println("Running:", strings.Join(cmdRun.Args, " "))
	cmdRun.Run()
}
