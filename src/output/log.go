package output

import "fmt"

func Info(s string) {
	fmt.Printf("\u001B[36mInfo: \u001B[0m%s\n", s)
}

func Error(s string) {
	fmt.Printf("\u001B[31mError: \u001B[0m%s\n", s)
}
