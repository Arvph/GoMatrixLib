package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

/**Вспомогательные функции**/

// GetElement предназначен для получения значения элемента матрицы по заданным индексам строки (rows) и столбца (cols).
func (m *Matrix) GetElement(rows, cols int) (float64, error) {
	if !m.isValid() {
		return 0, fmt.Errorf("invalid source matrix")
	} else if rows < 0 || cols < 0 || rows >= m.rows_ || cols >= m.cols_ {
		return 0, fmt.Errorf("row/col out of range")
	}
	return m.matrix_[rows][cols], nil
}

// Установка значения элемента матрицы по заданным индексам строки (rows) и столбца (cols)
func (m *Matrix) SetElement(rows, cols int, num float64) error {
	if !m.isValid() {
		return fmt.Errorf("invalid source matrix")
	} else if rows < 0 || cols < 0 || rows >= m.rows_ || cols >= m.cols_ {
		return fmt.Errorf("row/col out of range")
	}
	m.matrix_[rows][cols] = num
	return nil
}

// Печатает матрицу.  precision - количество знаков после запятой
func (m *Matrix) Print(precision int) {
	// p_nums := 1 // количестко знаков после точки
	if m.matrix_ == nil {
		fmt.Println("matrix: nil")
		return
	}

	maxColumnWidths := make([]int, len(m.matrix_[0]))
	for _, row := range m.matrix_ {
		for j, value := range row {

			width := func(value float64) int {
				return int(math.Log10(value)) + 1 + precision + 1
			}(value)

			if width > maxColumnWidths[j] {
				maxColumnWidths[j] = width
			}
		}
	}

	for i := range m.matrix_ {
		for j, value := range m.matrix_[i] {
			fmt.Printf("%*.*f  ", maxColumnWidths[j], precision, value)
		}
		fmt.Println()
	}
}

func (m *Matrix) IsEpmty() bool {
	return m.cols_ == 0 && m.rows_ == 0 && m.matrix_ == nil
}

func (m *Matrix) isValid() bool {
	return m.cols_ > 0 && m.rows_ > 0 && m.matrix_ != nil
}

// возвращает количество строк
func (m *Matrix) GetRows() int {
	return m.rows_
}

// возвращает количество столбцов
func (m *Matrix) GetCols() int {
	return m.cols_
}

// matrix возвращает внутреннюю матрицу для объекта Matrix
func (m *Matrix) matrix() [][]float64 {
	return m.matrix_
}

// метод заполняется матрицу случайными значениями
func (m *Matrix) Randomize() error {
	for i := 0; i < m.rows_; i++ {
		for j := 0; j < m.cols_; j++ {
			m.matrix_[i][j] = float64(rand.Intn(100) + 1)
		}
	}
	return nil
}

// функция вызывает метод Randomize, который заполняется матрицу случайными значениями
func RandomMatrix(other IMatrix) error {
	return other.Randomize()
}
