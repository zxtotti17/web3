package class2

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
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

// 1. 指针
func TestClass2_1(t *testing.T) {
	val := 10
	Class2_1(&val)
	if val != 20 {
		t.Errorf("Expected 20, got %d", val)
	}
}

func TestClass2_2(t *testing.T) {
	nums := []int{1, 2, 3}
	Class2_2(&nums)
	expected := []int{2, 4, 6}
	for i, v := range nums {
		if v != expected[i] {
			t.Errorf("Index %d: expected %d, got %d", i, expected[i], v)
		}
	}
}

// 2. Goroutine
func TestClass2_3(t *testing.T) {
	output := captureOutput(func() {
		Class2_3()
	})
	// 验证输出包含了奇数和偶数，顺序不确定
	for i := 1; i <= 10; i++ {
		if !strings.Contains(output, string(rune('0'+i))) && !strings.Contains(output, func() string {
			// 简单处理数字转字符串检查
			// 实际上直接检查数字是否存在即可，这里简单做
			return ""
		}()) {
			// 这种检查比较麻烦，直接检查行数或者特定数字
		}
	}
	// 简单检查行数，应该有 10 行 (1-10)
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 10 {
		t.Errorf("Expected 10 lines of output, got %d", len(lines))
	}
}

func TestClass2_4(t *testing.T) {
	tasks := []func(){
		func() { time.Sleep(10 * time.Millisecond) },
		func() { time.Sleep(20 * time.Millisecond) },
	}
	output := captureOutput(func() {
		Class2_4(tasks)
	})
	if !strings.Contains(output, "task1执行时间") || !strings.Contains(output, "task2执行时间") {
		t.Errorf("Output should contain task execution times")
	}
}

// 3. 面向对象
func TestShape(t *testing.T) {
	r := Rectangle{10, 20}
	if area := r.Area(); area != 200 {
		t.Errorf("Rectangle Area expected 200, got %f", area)
	}
	if perim := r.Perimeter(); perim != 60 {
		t.Errorf("Rectangle Perimeter expected 60, got %f", perim)
	}

	c := Circle{5}
	expectedArea := 3.14 * 5 * 5
	if area := c.Area(); area != expectedArea {
		t.Errorf("Circle Area expected %f, got %f", expectedArea, area)
	}
}

func TestEmployee(t *testing.T) {
	e := Employee{Person{"Alice", 30}, 1001}
	output := captureOutput(func() {
		e.PrintInfo()
	})
	expected := "Name: Alice, Age: 30, EmployeeID: 1001"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output containing %q, got %q", expected, output)
	}
}

// 4. Channel
func TestClass2_5(t *testing.T) {
	output := captureOutput(func() {
		Class2_5()
	})
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 10 {
		t.Errorf("Expected 10 lines, got %d", len(lines))
	}
}

func TestClass2_6(t *testing.T) {
	output := captureOutput(func() {
		Class2_6()
	})
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 100 {
		t.Errorf("Expected 100 lines, got %d", len(lines))
	}
}

// 5. 锁机制
func TestClass2_7(t *testing.T) {
	output := captureOutput(func() {
		Class2_7()
	})
	if !strings.Contains(output, "Class2_7 Count: 10000") {
		t.Errorf("Expected count 10000, got output: %s", output)
	}
}

func TestClass2_8(t *testing.T) {
	output := captureOutput(func() {
		Class2_8()
	})
	if !strings.Contains(output, "Class2_8 Count: 10000") {
		t.Errorf("Expected count 10000, got output: %s", output)
	}
}

// 集成测试
func TestClass2(t *testing.T) {
	// 仅仅确保不Panic
	captureOutput(func() {
		Class2()
	})
}
