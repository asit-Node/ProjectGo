package numberFrequency

// Get the frequecny of digits in the given string
func GetNumberFrequency(text string) (result []int) {
	slice := make([]int, 10) //  Default initialization to 0

	for inputIndex := 0; inputIndex < len(text); inputIndex++ {
		// Input in ascii 48 = 0 and 57 = 9
		if text[inputIndex] >= 48 && text[inputIndex] <= 57 {
			slice[text[inputIndex]-48]++
		}
	}

	result = slice
	return
}
