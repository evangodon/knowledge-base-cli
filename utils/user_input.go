package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput returns the user input
func GetUserInput(question string) string {
	fmt.Println(question)
	in := bufio.NewReader(os.Stdin)
	answer, err := in.ReadString('\n')

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(answer)
}

// ConfirmOperation asks the user to confirm the operation and exists if not
func ConfirmOperation(question string) {
	fmt.Printf("%s (y/n)", question)
	in := bufio.NewReader(os.Stdin)

	answer, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	answer = strings.TrimSpace(answer)
	if answer != "y" {
		os.Exit(0)
	}
}
