package main

import (
	"fmt"
	"github.com/cvstebut/hellogo/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("Hello, world. 3. Added to repo."))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(cmp.Diff("Hello World", "Hello World"))
}
