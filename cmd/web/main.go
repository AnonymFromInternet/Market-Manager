package main

import (
	"fmt"
	"github.com/anonymfrominternet/Market-Manager/pkg"
)

func main() {
	data := "abaccab"
	substringLength := pkg.HowLongIsSubstring(data)
	fmt.Println(fmt.Sprintf("Length of substring is %d", substringLength))
}

// Примеры:"abcdef" -> 2       "aaaaa" -> 5"        abaccab" -> 4

// result += ""
// for :range вроде можно со строкой -> if result не содержит substring -> result += ""
// Проблема в том, что, если буквы одинаковые то не вызовет ли это проблему?
// Также проблема в том, что цикл может остановиться еще до конца строки
