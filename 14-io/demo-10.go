package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		var txt = scanner.Text()
		if txt == "exit" {
			break
		}
		writer.WriteString(txt)
		writer.Flush()

	}
}
