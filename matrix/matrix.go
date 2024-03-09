// Библиотека, предназначенная для работы с матрицами.
package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

// Интерфейс IMatrix используется для создания общего контракта (набора методов),
// который должны реализовать структуры Matrix и SquareMatrix
type IMatrix interface {
	// CopyM(other IMatrix) error
	MoveM(other IMatrix) error
	Equal(other IMatrix) bool
	Sum(other IMatrix) (IMatrix, error)
	Subtract(other IMatrix) (IMatrix, error)
	Multiply(other IMatrix) (IMatrix, error)
	MultiplyByNumber(num float64) (IMatrix, error)
	GetElement(row, col int) (float64, error)
	// SetElement(row, col int, num float64) error
	GetRows() int
	GetCols() int
	Randomize() error
	Print(precision int)
	isValid() bool
}

// Основная структура
type Matrix struct {
	rows_   int
	cols_   int
	matrix_ [][]float64
}

// Квадратичная матрица. Реализация наследования от Matrix
type SquareMatrix struct {
	Matrix
}

/**Конструкторы и деструкторы**/
// Создаем новую матрицу и инициализируем нулями
func New(rows, cols int) (*Matrix, error) {
	if rows < 0 || cols < 0 {
		return nil, fmt.Errorf("invalid row or col size")
	}

	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	return &Matrix{
		rows_:   rows,
		cols_:   cols,
		matrix_: matrix,
	}, nil
}

// Создаем новую квадратную матрицу и инициализируем нулями
func NewSqr(size int) (*SquareMatrix, error) {
	if size < 0 {
		return nil, fmt.Errorf("invalid row or col size")
	}
	matrix := make([][]float64, size)
	for i := range matrix {
		matrix[i] = make([]float64, size)
	}

	return &SquareMatrix{
		Matrix: Matrix{
			rows_:   size,
			cols_:   size,
			matrix_: matrix}}, nil
}

// Конструктор копирования other в m
func (m *Matrix) Copy(other IMatrix) error {
	if !other.isValid() {
		return fmt.Errorf("invalid source matrix")
	}
	tmplMatrix, err := New(other.GetRows(), other.GetCols())
	if err != nil {
		return err
	}

	// Можно провести копирование срезов с использованием пакета reflect.
	// Стоит провести бенчмарк обоих подходов и опрделеить более эффективный подход
	for i := 0; i < other.GetRows(); i++ {
		for j := 0; j < other.GetCols(); j++ {
			tmpl, err := other.GetElement(i, j)
			if err != nil {
				return fmt.Errorf("out of range of mstrix")
			}
			tmplMatrix.matrix()[i][j] = tmpl
			// tmplMatrix.matrix()[i][j] = other.matrix()[i][j]
		}
	}
	m.cols_ = tmplMatrix.cols_
	m.rows_ = tmplMatrix.rows_
	m.matrix_ = tmplMatrix.matrix_
	return nil
}

// Конструктор переноса other в m
func (m *Matrix) MoveM(other IMatrix) error {
	o, ok := other.(*Matrix)
	if !ok {
		return fmt.Errorf("invalid source matrix")
	}
	if !o.isValid() {
		return fmt.Errorf("invalid source matrix")
	}
	m.RemoveM()

	m.cols_ = o.GetCols()
	m.rows_ = o.GetRows()
	m.matrix_ = o.matrix()

	o.RemoveM()
	return nil
}

// Условный деструктор - освобождает ресурсы, обнуляя данные в матрице и устанавливая размеры в 0.
func (m *Matrix) RemoveM() {
	for i := range m.matrix_ {
		for j := range m.matrix_[i] {
			m.matrix_[i][j] = 0.0
		}
	}
	m.rows_ = 0
	m.cols_ = 0
	m.matrix_ = nil
}

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
