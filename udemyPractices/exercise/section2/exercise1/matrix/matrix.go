package matrix

import (
	"errors"
	"fmt"
)

func New(rows ...[]float64) (Matrix, error) {
	if len(rows) == 0 {
		return Matrix{}, errors.New("no rows provided")
	}

	if len(rows[0]) == 0 {
		return Matrix{}, errors.New("no columns provided")
	}

	columns := len(rows[0])

	for i, row := range rows {
		if len(row) != columns {
			return Matrix{}, errors.New(fmt.Sprintf("row %d does not have the same number of columns", i))
		}
	}

	return Matrix{
		Values:   rows,
		Rows:     len(rows),
		Columns:  columns,
		IsSquare: len(rows) == columns,
	}, nil
}

type Matrix struct {
	Values   [][]float64
	Rows     int
	Columns  int
	IsSquare bool
}

func (m Matrix) Print() {
	for _, row := range m.Values {
		fmt.Println(row)
	}

	fmt.Printf("%d x %d\n", m.Rows, m.Columns)
	fmt.Printf("Cuadratic: %v\n", m.IsSquare)
}
