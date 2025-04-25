package tools

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func WaitForQ(ctx context.Context) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter command (q/quit/exit): ")
		scanner.Scan()
		fmt.Println(scanner.Text())

		if scanner.Err() != nil {
			fmt.Println("Error: ", scanner.Err())
		}

		if "q" == scanner.Text() || "quit" == scanner.Text() || "exit" == scanner.Text() {
			fmt.Println("Exiting gracefully ")
			os.Exit(0)
		}

	}
}
