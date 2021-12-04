package utils

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

// FzfSelect returns the selected line from the given data
func FzfSelect(data io.Reader) (string, error) {
	var result strings.Builder
	cmd := exec.Command("fzf", "--height", "40%")
	cmd.Stdout = &result
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()

	if err != nil {
		return "", err
	}

	_, err = io.Copy(stdin, data)
	if err != nil {
		return "", err
	}

	err = stdin.Close()
	if err != nil {
		return "", err
	}

	err = cmd.Start()
	if err != nil {
		return "", err
	}

	err = cmd.Wait()

	if err != nil {
		if err.Error() == "exit status 130" {
			os.Exit(0)
		}

		return "", err
	}

	return strings.TrimSpace(result.String()), nil
}
