package class2

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 值增加10
func Class2(p *int) {
	*p += 10
}

// 接收一个整数切片的指针，将切片中的每个元素乘以2
func Class2_2(p *[]int) {
	for i := range *p {
		(*p)[i] *= 2
	}
}

// go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
func Class2_3() {
	go func() {
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
	go func() {
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
func Class2_4(tasks []func()) {
	for num, task := range tasks {
		go func(task func()) {
			start := time.Now()
			task()
			end := time.Now()
			fmt.Println("task" + fmt.Sprintf("%d", num+1) + "执行时间:" + end.Sub(start).String())
		}(task)
	}
}

// 定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// 使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n", e.Name, e.Age, e.EmployeeID)
}

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来
func Class2_5() {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func Class2_6() {
	ch := make(chan int, 100)
	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}
}

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
func Class2_7() {
	var count int
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				count++
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Class2_7 Count:", count)
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func Class2_8() {
	var count int32
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("Class2_8 Count:", count)
}

func main() {
	calss2 := 1
	Class2(&calss2)
	fmt.Println("Class2:", calss2)

	nums := []int{1, 2, 3, 4, 5}
	Class2_2(&nums)
	fmt.Println("Class2_2:", nums)

	Class2_3()

	tasks := []func(){
		func() {
			time.Sleep(1 * time.Second)
		},
		func() {
			fmt.Println("你好啊")
		},
		func() {
			for i := 0; i < 1000000; i++ {
			}
		},
	}
	Class2_4(tasks)

	r := Rectangle{10, 20}
	c := Circle{5}
	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())
	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())

	e := Employee{Person{"Tom", 25}, 1001}
	e.PrintInfo()

	Class2_5()

	Class2_6()

	Class2_7()

	Class2_8()
}
