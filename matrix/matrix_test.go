package matrix

import (
	"testing"
)

func TestCopy_1(t *testing.T) {
	m1, _ := New(4, 3)
	m2, _ := New(4, 3)

	m1.Randomize()
	m2.Copy(m1)

	if !m1.Equal(m2) {
		t.Errorf("matrices must be equal")
	}
}

func TestCopy_2(t *testing.T) {
	m1, _ := New(4, 3)
	m2, _ := New(3, 4)

	m1.Randomize()

	if m2.Copy(m1) != nil {
		t.Errorf("matrices must be equal")
	}
}

func TestCopy_3(t *testing.T) {
	m1, _ := New(4, 3)
	m2, _ := NewSqr(3)

	m1.Randomize()

	if m2.Copy(m1) != nil {
		t.Errorf("matrices must be equal")
	}
}

func TestCopy_4(t *testing.T) {
	m1, _ := NewSqr(3)
	m2, _ := New(4, 3)

	m1.Randomize()

	if m2.Copy(m1) != nil {
		t.Errorf("matrices must be equal")
	}
}

func TestCopy_5(t *testing.T) {
	m1, _ := NewSqr(3)
	m2, _ := NewSqr(3)

	m1.Randomize()
	m2.Copy(m1)

	if !m2.Equal(m1) {
		t.Errorf("matrices must be equal")
	}
}

func TestMove_1(t *testing.T) {
	m1, _ := NewSqr(3)
	m2, _ := NewSqr(0)

	m1.Randomize()
	m2.Move(m1)

	if !m1.IsEpmty() || m2.IsEpmty() {
		t.Errorf("matrices must be equal")
	}
}

func TestMove_2(t *testing.T) {
	m1, _ := New(3, 5)
	m2, _ := NewSqr(0)

	m1.Randomize()
	m2.Move(m1)

	if !m1.IsEpmty() || m2.IsEpmty() {
		t.Errorf("matrices must be equal")
	}
}

func TestMove_3(t *testing.T) {
	m1, _ := NewSqr(3)
	m2, _ := New(1, 5)

	m1.Randomize()
	m2.Move(m1)

	if !m1.IsEpmty() || m2.IsEpmty() {
		t.Errorf("matrices must be equal")
	}
}

func TestMove_4(t *testing.T) {
	m1, _ := NewSqr(7)
	m2, _ := NewSqr(3)

	m1.Randomize()
	m2.Move(m1)

	if !m1.IsEpmty() || m2.IsEpmty() {
		t.Errorf("matrices must be equal")
	}
}

func TestMove_5(t *testing.T) {
	m1, _ := NewSqr(0)
	m2, _ := New(1, 5)

	m1.Randomize()
	m2.Move(m1)

	if m1.IsEpmty() || m2.IsEpmty() {
		t.Errorf("course matrix must not be nil")
	}
}

func TestEqual_1(t *testing.T) {
	m1, _ := New(3, 4)
	m2, _ := New(4, 3)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_2(t *testing.T) {
	m1, _ := New(4, 4)
	m2, _ := New(4, 4)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_3(t *testing.T) {
	m1, _ := New(4, 4)
	m2, _ := New(4, 4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}

func TestEqual_4(t *testing.T) {
	m1, _ := NewSqr(4)
	m2, _ := NewSqr(3)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_5(t *testing.T) {
	m1, _ := NewSqr(4)
	m2, _ := NewSqr(4)
	m1.Randomize()
	m2.Randomize()

	res := m1.Equal(m2)
	if res == true {
		t.Errorf("matrices must not be equal")
	}
}

func TestEqual_6(t *testing.T) {
	m1, _ := NewSqr(4)
	m2, _ := NewSqr(4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}

func TestEqual_7(t *testing.T) {
	m1, _ := New(4, 4)
	m2, _ := NewSqr(4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}

func TestEqual_8(t *testing.T) {
	m1, _ := NewSqr(4)
	m2, _ := New(4, 4)
	m1.Randomize()
	m1.Copy(m2)

	res := m1.Equal(m2)
	if res == false {
		t.Errorf("matrices must be equal")
	}
}
