package scrabble

import (
	"strings"
)

func Score(word string) int {
  upperCaseArray := strings.Split(strings.ToUpper(word), "")

  sumScore := 0

  for _, e := range upperCaseArray{

    switch e {
    case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":  sumScore += 1
    case "D", "G":  sumScore += 2
    case "B", "C", "M", "P":  sumScore += 3
    case "F", "H", "V", "W", "Y":  sumScore += 4
    case "K":  sumScore += 5
    case "J", "X":  sumScore += 8
    case "Q", "Z":  sumScore += 10
    }
  }
  return sumScore
}
