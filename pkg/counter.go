package pkg

func HowLongIsSubstring(data string) int {
	var letters = make(map[string]int)
	var prevLetter string

	var length int

	for i := 0; i < len(data); i++ {
		if len(letters) < 2 {
			if _, ok := letters[string(data[i])]; !ok {
				letters[string(data[i])] += 1
				prevLetter = string(data[i])
				length = 0
				for _, value := range letters {
					length += value
				}
			} else {
				letters[string(data[i])] += 1
				prevLetter = string(data[i])
				length = 0
				for _, value := range letters {
					length += value
				}
			}
		} else {
			if _, ok := letters[string(data[i])]; !ok {
				length = 0
				for _, value := range letters {
					length += value
				}
				letters = make(map[string]int)
				letters[prevLetter] = 1
				letters[string(data[i])] += 1
				prevLetter = string(data[i])
			} else {
				length = 0
				for _, value := range letters {
					length += value
				}
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
