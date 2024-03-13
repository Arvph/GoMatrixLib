// Библиотека, предназначенная для работы с матрицами.
package matrix

import "fmt"

// Интерфейс IMatrix используется для создания общего контракта (набора методов),
// который должны реализовать структуры Matrix и SquareMatrix
type IMatrix interface {
	IMatrixOperations
	IMatrixElementAccess
	IMatrixProperties
	IMatrixUtility
}

// Интерфейс для основных матричных операций
type IMatrixOperations interface {
	Remove()
	Move(other IMatrix) error
	Equal(other IMatrix) bool
	Sum(other IMatrix) (IMatrix, error)
	Subtract(other IMatrix) (IMatrix, error)
	Multiply(other IMatrix) (IMatrix, error)
	MultiplyByNumber(num float64) (IMatrix, error)
}

// Интерфейс для доступа к элементам матрицы
type IMatrixElementAccess interface {
	GetElement(row, col int) (float64, error)
	GetRows() int
	GetCols() int
}

// Интерфейс для получения свойств матрицы
type IMatrixProperties interface {
	matrix() [][]float64
	isValid() bool
	IsEpmty() bool
}

// Интерфейс для вспомогательных функций
type IMatrixUtility interface {
	Randomize() error
	Print(precision int)
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

	for i := 0; i < other.GetRows(); i++ {
		for j := 0; j < other.GetCols(); j++ {
			tmpl, err := other.GetElement(i, j)
			if err != nil {
				return fmt.Errorf("out of range of mstrix")
			}
			tmplMatrix.matrix()[i][j] = tmpl
		}
	}
	m.cols_ = tmplMatrix.cols_
	m.rows_ = tmplMatrix.rows_
	m.matrix_ = tmplMatrix.matrix_
	return nil
}

// Конструктор переноса other в m
func (m *Matrix) Move(other IMatrix) error {
	if !other.isValid() {
		return fmt.Errorf("invalid source matrix")
	}
	m.Remove()

	if err := m.Copy(other); err != nil {
		return err
	}

	other.Remove()
	return nil
}

// Условный деструктор - освобождает ресурсы, обнуляя данные в матрице и устанавливая размеры в 0.
func (m *Matrix) Remove() {
	for i := range m.matrix_ {
		for j := range m.matrix_[i] {
			m.matrix_[i][j] = 0.0
		}
	}
	m.rows_ = 0
	m.cols_ = 0
	m.matrix_ = nil
}
