package php8

import "slices"

type NewLines struct {
	data []int
}

func (newLines *NewLines) Append(position int) {
	if len(newLines.data) == 0 || newLines.data[len(newLines.data)-1] < position {
		newLines.data = append(newLines.data, position)
	}
}

func (newLines *NewLines) GetLine(position int) int {
	line := len(newLines.data) + 1

	for i, newLinePosition := range slices.Backward(newLines.data) {
		if position < newLinePosition {
			line = i + 1
		} else {
			break
		}
	}

	return line
}
