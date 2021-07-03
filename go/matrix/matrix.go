package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type matrix struct {
	rows    [][]int
	columns [][]int
}
type Matrix interface {
	Rows() [][]int
	Cols() [][]int
	Set(row int, col int, val int) bool
}

func (m matrix) Rows() [][]int {
	cpy := make([][]int, len(m.rows))
	for i, col := range m.rows {
		cpy[i] = make([]int, len(col))
		copy(cpy[i], col)
	}
	return cpy
}

func (m matrix) Cols() [][]int {
	cpy := make([][]int, len(m.columns))
	for i, row := range m.columns {
		cpy[i] = make([]int, len(row))
		copy(cpy[i], row)
	}
	return cpy
}
func (m matrix) Set(row int, col int, val int) bool {
	if row >= len(m.rows) || col >= len(m.columns) || row < 0 || col < 0 {
		return false
	}
	m.rows[row][col] = val
	m.columns[col][row] = val
	return true
}

func New(s string) (Matrix, error) {
	if s == "" {
		return nil, errors.New("string can't be empty")
	}

	inputLines := strings.Split(s, "\n")
	//var inputBlocks [][]int
	var rows [][]int
	var cols [][]int

	numOfRows := len(inputLines)
	for indexrow, row := range inputLines {
		if row == "" {
			return nil, errors.New("a row was empty")
		}
		rowSplit := filter(strings.Split(row, " "), func(i string) bool {
			if i != "" {
				return true
			} else {
				return false
			}
		})
		numOfCols := len(rowSplit)
		for indexcol, val := range rowSplit {

			if rows == nil {
				rows = make([][]int, numOfRows)
				for i := 0; i < numOfRows; i++ {
					rows[i] = make([]int, numOfCols)
				}
			}

			if cols == nil {
				cols = make([][]int, numOfCols)
				for i := 0; i < numOfCols; i++ {
					cols[i] = make([]int, numOfRows)
				}
			}

			if numOfCols != len(rows[0]) {
				return nil, errors.New("not every row has the same amount of columns")
			}

			if val == "" {
				continue
			}

			val, err := strconv.Atoi(val)
			rows[indexrow][indexcol] = val
			cols[indexcol][indexrow] = val
			if err != nil {
				return nil, err
			}
		}
	}
	matr := matrix{
		rows:    rows,
		columns: cols,
	}
	return matr, nil
}
