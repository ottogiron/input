# input
Simple user input from stdin  go library. Will loop until a valid value is entered.

## Available methods

* input.GetString(label string)  string
* input.GetInt32(label string)   int32
* input.GetInt64(label string)   int64
* input.GetFloat64(label string) float64
* input.GetFloat32(label string) float32

## Example

```go
package main

import (
	"fmt"

	"github.com/ottogiron/input"
)

func main() {

	s := input.GetString("String Value:")
	fmt.Println("value", s)
	i32 := input.GetInt32("int32 value")
	fmt.Println("value", i32)
	i64 := input.GetInt64("int64 value")
	fmt.Println("value: ", i64)
	f32 := input.GetFloat32("float32 value")
	fmt.Println("value:", f32)
	f64 := input.GetFloat64("float 64:")
	fmt.Println("value:", f64)

}

```

### Output

```bash
String Value: kjfkjfkdjf
value kjfkjfkdjf
int32 value kdjfkdf
int32 value kdjfkdjf
int32 value 2
value 2
int64 value kdjfkjdkjf
int64 value 434
value:  434
float32 value 3
value: 3
float 64: 3.42
value: 3.42
```
