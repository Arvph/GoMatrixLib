package matrix_test

import (
	"testing"

	"github.com/arvph/GoMatrixLib/matrix"
)

func TestEqual_1(t *testing.T) {
	m1, _ := matrix.New(3, 4)
	m2, _ := matrix.New(4, 3)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_2(t *testing.T) {
	m1, _ := matrix.New(4, 4)
	m2, _ := matrix.New(4, 4)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_3(t *testing.T) {
	m1, _ := matrix.New(4, 4)
	m2, _ := matrix.New(4, 4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}

func TestEqual_4(t *testing.T) {
	m1, _ := matrix.NewSqr(4)
	m2, _ := matrix.NewSqr(3)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_5(t *testing.T) {
	m1, _ := matrix.NewSqr(4)
	m2, _ := matrix.NewSqr(4)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_6(t *testing.T) {
	m1, _ := matrix.NewSqr(4)
	m2, _ := matrix.NewSqr(4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}

func TestEqual_7(t *testing.T) {
	m1, _ := matrix.New(4, 4)
	m2, _ := matrix.NewSqr(4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}
