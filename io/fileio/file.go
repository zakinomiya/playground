package fileio

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Stack interface {
	Push(i string) []string
	Pop() *string
	Size() int
}

type StringStack struct {
	data []string
}

func New() *StringStack {
	return &StringStack{
		data: make([]string, 0, 0),
	}
}

func (ss *StringStack) Push(i string) []string {
	ss.data = append(ss.data, i)
	return ss.data
}

func (ss *StringStack) Pop() *string {
	if ss.Size() == 0 {
		return nil
	}
	// the index of the last element in the stack
	n := len(ss.data) - 1
	i := ss.data[n]
	ss.data = ss.data[:n]

	return &i
}

func (ss *StringStack) Size() int {
	return len(ss.data)
}

func (ss *StringStack) Reverse() {
	reversed := New()

	for ss.Size() > 0 {
		if i := ss.Pop(); i != nil {
			reversed.Push(*i)
		}
	}

	ss.data = reversed.data
}

func ReadFileAndPrintReverse(fp *os.File) {
	sc := bufio.NewScanner(fp)
	stack := New()
	for sc.Scan() {
		stack.Push(sc.Text())
	}

	if err := sc.Err(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}

	// implementation 1: reversed and join string array
	stack.Reverse()
	fmt.Print(strings.Join(stack.data, "\n"))

	// implementation 2: loop and print
	//for stack.Size() > 0 {
	//  fmt.Print(*stack.Pop())
	//}
}
