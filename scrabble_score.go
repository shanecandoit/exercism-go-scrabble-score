package scrabble

import "strings"

// Score the scrabble score of a word
func Score(word string) (score int) {

	score = 0
	word = strings.ToUpper(word)

	// dont index into string because unicode, even though values of concern are ascii
	for _, ch := range word {
		switch ch {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}

	return score
}
