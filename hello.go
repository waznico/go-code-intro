package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/waznico/hello/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello world", "Hello Go"))
}
