package numberFrequency

import (
	"errors"
	//"fmt"
	"testing"
)

func TestGetNumberFrequency(t *testing.T) {
	testString := "12378909"
	ExpectedResult := []int{1, 1, 1, 1, 0, 0, 0, 1, 1, 2} // Expected result for the testString.
	result := GetNumberFrequency(testString)
	if len(result) == 0 {
		t.Fatalf("No results were returned", errors.New("No content"))
	}
	for resultIndex := 0; resultIndex < len(result); resultIndex++ {
		//fmt.Println("Results ", result[resultIndex], ExpectedResult[resultIndex])
		if result[resultIndex] != ExpectedResult[resultIndex] {
			t.Fatalf("Incorrect result returned. Expected %d  Actual %d at index %d", result[resultIndex], ExpectedResult[resultIndex], resultIndex, errors.New("Mismatch"))
		}
	}
}
