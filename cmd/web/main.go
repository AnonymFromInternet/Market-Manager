package main

import (
	"fmt"
	"github.com/anonymfrominternet/Market-Manager/pkg"
)

func main() {
	data := "abaccab"
	substringLength := pkg.HowLongIsSubstring(data)
	fmt.Println(fmt.Sprintf("Length of the substring is %d", substringLength))
}
