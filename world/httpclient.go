package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	buf, err := json.Marshal(nil)
	fmt.Println(string(buf), err)
}
