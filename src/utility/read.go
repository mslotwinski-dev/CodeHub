package utility

import (
	"bufio"
	"strings"
)

func Read(reader *bufio.Reader) string {

	v, _ := reader.ReadString('\n')
	v = strings.TrimSpace(v)

	return v
}
