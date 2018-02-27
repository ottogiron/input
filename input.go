package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

//GetString asks the user for a string
func GetString(label string) string {
	for {
		fmt.Print(label, " ")
		val, err := reader.ReadString('\n')
		if err == nil {
			return strings.TrimSpace(val)
		}
	}
}
