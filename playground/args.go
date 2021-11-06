// Demonstrate the usage of os.Args
// Also - strings.Join
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args)

	joined := strings.Join(os.Args, " - ")
	fmt.Println(joined)
}
