package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func Fzf(data io.Reader) (string, error) {
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

	fmt.Println("got here")
	return strings.TrimSpace(result.String()), nil

}
