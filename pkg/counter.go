package pkg

func HowLongIsSubstring(data string) int {
	var letters = make(map[string]int)
	var prevLetter string

	var length int

	for i := 0; i < len(data); i++ {
		if len(letters) < 2 {
			letters[string(data[i])] += 1
			prevLetter = string(data[i])
			lengthRefresh(&length, letters)
		} else {
			if _, ok := letters[string(data[i])]; !ok {
				lengthRefresh(&length, letters)
				letters = make(map[string]int)
				letters[prevLetter] = 1
				letters[string(data[i])] += 1
				prevLetter = string(data[i])
			} else {
				lengthRefresh(&length, letters)
				letters[string(data[i])] += 1
				prevLetter = string(data[i])
			}
		}
	}
	return length
}

func lengthRefresh(length *int, letters map[string]int) int {
	*length = 0
	for _, value := range letters {
		*length += value
	}
	return *length
}
