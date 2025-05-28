package internal

import (
	"bufio"
	"os"
	"strings"
)

func LoadCommands(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var commands []string
	var inCommands bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "[Commands]" {
			inCommands = true
			continue
		}
		if inCommands {
			if strings.HasPrefix(line, "[") { // 別のセクションに入ったら終了
				break
			}
			if line != "" {
				commands = append(commands, line)
			}
		}
	}
	return commands, scanner.Err()
}
