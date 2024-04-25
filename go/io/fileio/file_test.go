package fileio

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var testTargetLines = []int{1000, 100000, 1000000}

func prepareFiles(lines ...int) {
	for _, l := range lines {
		filedata := make([]string, l, l)
		for i := 0; i < l; i++ {
			filedata[i] = fmt.Sprint(i)
		}

		if err := os.WriteFile(fmt.Sprint(l), []byte(strings.Join(filedata, "\n")), 0444); err != nil {
			if err != os.ErrExist {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
	}
}

func rmFiles(lines ...int) {
	for _, v := range lines {
		os.Remove(fmt.Sprint(v))
	}
}

func TestMain(m *testing.M) {
	prepareFiles(testTargetLines...)

	exitCode := m.Run()

	rmFiles(testTargetLines...)
	os.Exit(exitCode)
}

func BenchmarkReadFileAndPrintReverse_1000(t *testing.B) {
	fp, _ := os.Open(fmt.Sprint(testTargetLines[0]))
	t.ResetTimer()
	ReadFileAndPrintReverse(fp)
}

func BenchmarkReadFileAndPrintReverse_100000(t *testing.B) {
	fp, _ := os.Open(fmt.Sprint(testTargetLines[1]))
	t.ResetTimer()
	ReadFileAndPrintReverse(fp)
}

func BenchmarkReadFileAndPrintReverse_1000000(t *testing.B) {
	fp, _ := os.Open(fmt.Sprint(testTargetLines[2]))
	t.ResetTimer()
	ReadFileAndPrintReverse(fp)
}
