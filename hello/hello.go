package main

import (
	"fmt"
	"time"

	"golang.org/x/advance/stringutil"
)

func main() {
	fmt.Println(fmt.Sprintf("%v", true))
	result := stringutil.Reverse("Hello Workspace")
	fmt.Println(result)
	timestrap := fmt.Sprintf("%d", time.Now().UnixMicro())
	fmt.Println(timestrap, len(timestrap))
	timestrap = fmt.Sprintf("%d", time.Now().Unix())
	fmt.Println(timestrap, len(timestrap))
	timestrap = fmt.Sprintf("%d", time.Now().UnixMilli())
	fmt.Println(timestrap, len(timestrap))
	timestrap = fmt.Sprintf("%d", time.Now().UnixNano())
	fmt.Println(timestrap, len(timestrap))
	timestrap = fmt.Sprintf("%d", time.Now().UnixNano()/1e6)
	fmt.Println(timestrap, len(timestrap))

}
