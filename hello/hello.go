package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	result := stringutil.Reverse("Hello Workspace")
	fmt.Println(result)
}
