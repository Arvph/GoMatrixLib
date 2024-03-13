package matrix

import "fmt"

/**МАТЕМАТИЧЕСКИЕ ФУНКЦИИ**/

// Equal сравнивает матрицу с другой матрицей, реализующей интерфейс IMatrix
func (m *Matrix) Equal(other IMatrix) bool {
	if m == nil || other == nil {
		return false
	}

	if m.rows_ != other.GetRows() || m.cols_ != other.GetCols() {
		return false
	}

	for i := 0; i < m.rows_; i++ {
		for j := 0; j < m.cols_; j++ {
			tmpl, err := other.GetElement(i, j)
			if err != nil {
				return false
			}
			if m.matrix_[i][j] != tmpl {
				return false
			}
		}
	}
	return true
}

// SumM суммирует матрицу с другой матрицей, реализующей интерфейс IMatrix, и возвращает результ
func (m *Matrix) Sum(other IMatrix) (IMatrix, error) {
	if m == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	o, ok := other.(*Matrix)
	if !ok {
		return nil, fmt.Errorf("invalid source matrix type")
	}

	if o == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	if m.rows_ != o.rows_ || m.cols_ != o.cols_ {
		return nil, fmt.Errorf("matrices have different sizes")
	}

	result := make([][]float64, m.rows_)
	for i := range result {
		result[i] = make([]float64, m.cols_)
		for j := 0; j < m.cols_; j++ {
			result[i][j] = m.matrix_[i][j] + o.matrix_[i][j]
		}
	}

	return &Matrix{
		rows_:   m.rows_,
		cols_:   m.cols_,
		matrix_: result,
	}, nil
}

// Subtract вычитает из матрицы другую матрицу, реализующую интерфейс IMatrix, и возвращает результ
func (m *Matrix) Subtract(other IMatrix) (IMatrix, error) {
	if m == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	o, ok := other.(*Matrix)
	if !ok {
		return nil, fmt.Errorf("invalid source matrix type")
	}

	if o == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	if m.rows_ != o.rows_ || m.cols_ != o.cols_ {
		return nil, fmt.Errorf("matrices have different sizes")
	}

	result := make([][]float64, m.rows_)
	for i := range result {
		result[i] = make([]float64, m.cols_)
		for j := 0; j < m.cols_; j++ {
			result[i][j] = m.matrix_[i][j] - o.matrix_[i][j]
		}
	}

	return &Matrix{
		rows_:   m.rows_,
		cols_:   m.cols_,
		matrix_: result,
	}, nil
}

// Multiply умножает матрицу на другую матрицу, реализующую интерфейс IMatrix, и возвращает результат умножения
func (m *Matrix) Multiply(other IMatrix) (IMatrix, error) {
	if m == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	o, ok := other.(*Matrix)
	if !ok {
		return nil, fmt.Errorf("invalid source matrix type")
	}
	if o == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	if m.cols_ != o.rows_ {
		return nil, fmt.Errorf("matrices sizes do not allow multiplication")
	}

	// Вычисление умножения матриц
	result := make([][]float64, m.rows_)
	for i := range result {
		result[i] = make([]float64, o.cols_)
		for j := 0; j < o.cols_; j++ {
			for k := 0; k < m.cols_; k++ {
				result[i][j] += m.matrix_[i][k] * o.matrix_[k][j]
			}
		}
	}

	return &Matrix{
		rows_:   m.rows_,
		cols_:   m.cols_,
		matrix_: result,
	}, nil
}

// MultiplyByNumber умножает матрицу на число и записывает результат в вызывающий объект
func (m *Matrix) MultiplyByNumber(num float64) (IMatrix, error) {
	if m == nil {
		return nil, fmt.Errorf("invalid source matrix")
	}

	result := make([][]float64, m.rows_)
	for i := range result {
		result[i] = make([]float64, m.cols_)
		for j := 0; j < m.cols_; j++ {
			result[i][j] = m.matrix_[i][j] * num
		}
	}

	return &Matrix{
		rows_:   m.rows_,
		cols_:   m.cols_,
		matrix_: result,
	}, nil
}
