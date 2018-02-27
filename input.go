package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

//GetInt32 asks the user for an int32 value
func GetInt32(label string) int32 {
	for {
		fmt.Print(label, " ")
		val, err := reader.ReadString('\n')
		if err == nil {
			val = strings.TrimSpace(val)
			i, err := strconv.ParseInt(val, 10, 32)
			if err == nil {
				return int32(i)
			}
		}
	}
}

//GetInt64 asks the user for an int32 value
func GetInt64(label string) int64 {
	for {
		fmt.Print(label, " ")
		val, err := reader.ReadString('\n')
		if err == nil {
			val = strings.TrimSpace(val)
			i, err := strconv.ParseInt(val, 10, 64)
			if err == nil {
				return i
			}
		}
	}
}

//GetFloat32 asks the user for an float32 value
func GetFloat32(label string) float32 {
	for {
		fmt.Print(label, " ")
		val, err := reader.ReadString('\n')
		if err == nil {
			val = strings.TrimSpace(val)
			f, err := strconv.ParseFloat(val, 32)
			if err == nil {
				return float32(f)
			}
		}
	}
}

//GetFloat64 asks the user for an float32 value
func GetFloat64(label string) float64 {
	for {
		fmt.Print(label, " ")
		val, err := reader.ReadString('\n')
		if err == nil {
			val = strings.TrimSpace(val)
			f, err := strconv.ParseFloat(val, 64)
			if err == nil {
				return f
			}
		}
	}
}
