package pascal

func Triangle(n int) (retval [][]int) {
	retval = make([][]int, n)
	for row := 0; row < n; row++ {
		retval[row] = make([]int, row+1)
		maxColumn := row
		column := 0
		for column < maxColumn+1 {
			if column == 0 || column == maxColumn {
				retval[row][column] = 1
			} else {
				retval[row][column] = retval[row-1][column-1] + retval[row-1][column]
			}
			column++
		}
	}
	return
}
