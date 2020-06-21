package main

import (
	"fmt"

	"github.com/boseabhishek/go-bazel-project/second"
)

func main() {
	output := second.LearnBazel("Abhi")
	fmt.Println(output)
}
