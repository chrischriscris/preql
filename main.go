package main

import (
	"fmt"
	"os"
)

func main() {
    args := os.Args
    if len(args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: preql <input_file>")
        return
    }


}
