package array_stack

import "testing"

func TestArrayStack(t *testing.T) {
	as := &ArrayStack{
		n:   3,
		cap: 3,
		buf: make([]int, 3),
	}

	// 0 : cap = 3
	as.Set(0, 0)
	if as.Get(0) != 0 {
		t.Fatal("cannot get correct value")
	}

	// 0, 1 : cap = 3
	as.Set(1, 1)
	// 0, 1, 2 : cap = 3
	as.Set(2, 2)
	if as.Size() != 3 {
		t.Fatal("now the len should 3")
	}

	// 0, 10, 2 : cap = 3
	if v := as.Set(1, 10); v != 1 {
		t.Fatalf("Set should return the original value at 1 which is 1, but actual: %d\n", v)
	}
	// 0, 10, 2 : cap = 3
	if v := as.Get(1); v != 10 {
		t.Fatalf("now the value at index 1 should be 10, but actual %d\n ", v)
	}

	// 0, 1, 2 : cap = 3
	as.Set(1, 1)

	// 0, 10, 1, 2 : cap = 6
	as.Add(1, 10)
	if as.Size() != 4 {
		t.Log(as)
		t.Fatal("size should be 4 now")
	}
	if as.Cap() != 6 {
		t.Log(as.buf)
		t.Fatal("cap should be 6 now")
	}
	if v := as.Get(1); v != 10 {
		t.Log(as.buf)
		t.Fatal("now the value at index 1 should be 10")
	}
	if v := as.Get(3); v != 2 {
		t.Log(as.buf)
		t.Fatal("now the value at index 3 should be 2")
	}

	// 0, 10, 1: cap = 6
	if v := as.Remove(3); v != 2 {
		t.Fatalf("Remove should return 2, but actual %d\n", v)
		t.Fatalf("%+v\n", as.buf)
	}

	// 0, 10: cap = 4
	as.Remove(2)
	if v := as.Cap(); v != 4 {
		t.Fatalf("now the cap should be 4 but actual %d\n", v)
	}
}
