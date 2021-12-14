package shared

type Matrix09 [][]int

func ParseMatrix09(lines []string) Matrix09 {
	matrix := Matrix09{}
	for _, line := range lines {
		row := []int{}
		for _, c := range line {
			row = append(row, int(c-'0'))
		}
		matrix = append(matrix, row)
	}
	return matrix
}
