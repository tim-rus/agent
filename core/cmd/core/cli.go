package main

import (
	"bufio"
	"fmt"
	"os"
)

func cliLoop() error {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("> ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		println("ECHO:", input)

	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
