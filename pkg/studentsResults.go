package pkg

type StudentsResults struct {
	StudentsAnswers [][]string
	CorrectAnswers  []string
}

func (studentsResults *StudentsResults) GetStudentsResults() int {
	resultsInFloat := make([]float32, len(studentsResults.StudentsAnswers))
	var sum float32
	studentsAmount := len(studentsResults.StudentsAnswers)

	for i := 0; i < len(studentsResults.StudentsAnswers); i++ {
		for j := 0; j < len(studentsResults.StudentsAnswers[i]); j++ {
			if studentsResults.StudentsAnswers[i][j] == studentsResults.CorrectAnswers[j] {

				resultsInFloat[i] += 100 / float32(len(studentsResults.CorrectAnswers))
			}
		}
	}

	for i := 0; i < len(resultsInFloat); i++ {
		sum += resultsInFloat[i]
	}
	return int(sum) / studentsAmount
}
