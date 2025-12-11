package class3

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

// 辅助函数：捕获 stdout 输出
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestClass3_1(t *testing.T) {
	output := captureOutput(func() {
		Class3_1()
	})
	fmt.Println(output)
}

func TestClass3_2(t *testing.T) {
	output := captureOutput(func() {
		Class3_2()
	})
	fmt.Println(output)
}
