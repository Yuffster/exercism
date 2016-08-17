package pascal

// Triangle ugh math.
func Triangle(n int) [][]int {
	rows := make([][]int, n)
	rows[0] = []int{1}
	rows[0][0] = 1
	for i := 1; i < n; i++ {
		rows[i] = make([]int, i+1)
		rows[i][0] = 1
		for j := 1; j < i; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i][i] = 1
	}
	return rows
}