package internal

import (
	"bytes"
	"os"
	"os/exec"
)

func SelectWithPeco(options []string) (string, error) {
	cmd := exec.Command("peco")
	cmd.Stdin = bytes.NewBufferString(joinLines(options))
	cmd.Stderr = os.Stderr

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(out.Bytes())), nil
}

func joinLines(list []string) string {
	return string(bytes.Join(func() [][]byte {
		b := make([][]byte, len(list))
		for i, s := range list {
			b[i] = []byte(s)
		}
		return b
	}(), []byte("\n")))
}

// 標準入出力を使いまわすユーティリティ
func StdStreams() (*os.File, *os.File, *os.File) {
	return os.Stdin, os.Stdout, os.Stderr
}
