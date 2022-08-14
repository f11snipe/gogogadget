package main

import (
	"fmt"

	"github.com/c2h5oh/datasize"
)

func main() {
	fmt.Printf("%s", (4096 * datasize.KB).String())
}
