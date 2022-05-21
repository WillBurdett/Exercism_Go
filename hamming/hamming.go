package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	aArray := strings.Split(a, "")
	bArray := strings.Split(b, "")

	if len(aArray) != len(bArray){return 0, errors.New("Input strings are of different lengths")}

	differences := 0
	for i, e := range aArray{
		if e != bArray[i]{
			differences++
		}
	}
	return differences, nil
}
