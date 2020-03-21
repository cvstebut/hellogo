package main

import (
	"fmt"
	"github.com/cvstebut/hellogo/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world. Two. Added to repo."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
