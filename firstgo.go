package main

import (
	"fmt"

	"github.com/boseabhishek/go-bazel-project/second"
)

func main() {
	output := second.AddSmile("Abhi")
	fmt.Println(output)
}
