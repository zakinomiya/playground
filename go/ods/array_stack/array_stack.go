package array_stack

type ArrayStack struct {
	n   int
	cap int
	buf []int
}

func (as *ArrayStack) Size() int {
	return as.n
}

func (as *ArrayStack) Cap() int {
	return as.cap
}

func (as *ArrayStack) Get(i int) int {
	return as.buf[i]
}

func (as *ArrayStack) Set(i int, x int) int {
	y := as.buf[i]
	as.buf[i] = x
	return y
}

func (as *ArrayStack) Add(i int, x int) {
	if as.Size() >= as.Cap() {
		as.Resize()
	}

	// replace elements by 1 right to left
	for j := as.Size(); j > i; j-- {
		as.buf[j] = as.buf[j-1]
	}

	as.buf[i] = x
	as.n++
}

func (as *ArrayStack) Remove(i int) int {
	elem := as.buf[i]
	for ; i < as.Size(); i++ {
		as.buf[i+1] = as.buf[i]
	}

	as.n--

	// resize if the cap is too large
	if as.Cap() >= as.Size()*3 {
		as.Resize()
	}

	return elem
}

func (as *ArrayStack) Resize() {
	as.cap = as.Size() * 2
	newBuf := make([]int, as.Cap())
	for i := 0; i < as.n; i++ {
		newBuf[i] = as.buf[i]
	}
	as.buf = newBuf
}
