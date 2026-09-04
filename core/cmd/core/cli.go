package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	msgWelcome = "hey there!"
	msgGoodbye = "bye"

	cmdExit = ":q"
)

func chat() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msgWelcome)
	fmt.Printf("[type '%s' to quit]\n", cmdExit)

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if strings.TrimSpace(input) == cmdExit {
			fmt.Println(msgGoodbye)
			break
		}

		fmt.Printf("you typed [%s]\n", input)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
