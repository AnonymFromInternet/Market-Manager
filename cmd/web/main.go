package main

import (
	"fmt"
	"github.com/anonymfrominternet/Market-Manager/pkg"
)

func main() {
	data := "abaccab"
	substringLength := pkg.HowLongIsSubstring(data)
	fmt.Println(fmt.Sprintf("Length of the substring is %d", substringLength))

	// Students Section
	studentsAnswers := [][]string{{"A", "C", "E", "D"}, {"E", "E", "E", "D"}, {"D", "C", "E", "D"}, {"A", "C", "D", "D"}}
	correctAnswers := []string{"A", "C", "E", "D"}

	answers := pkg.StudentsResults{
		StudentsAnswers: studentsAnswers,
		CorrectAnswers:  correctAnswers,
	}
	result := answers.GetStudentsResults()
	fmt.Println(fmt.Sprintf("Middle result is %d%%:", result))
}
