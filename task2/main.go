package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 指针
// 题目1
func pointInt(i *int) {
	*i += 10
}

// 题目2
func pointSlice(s *[]int) {
	_s := *s
	for v := range _s {
		_s[v] *= 2
	}
}

// Goroutine
// 题目1
func Go1to10() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
	}
}

func Go2to10() {
	for i := 2; i < 10; i += 2 {
		fmt.Println(i)
	}
}

// 题目2
func goRoutineTask(Task func()) {

	go func() {
		startTime := time.Now()
		Task()
		endTime := time.Since(startTime)
		fmt.Println(endTime)
	}()
}

// 面向对象
// 题目1
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}
type Circle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("Rectangle Area")
}

func (r *Rectangle) Perimeter() {
	fmt.Println("Rectangle Perimeter")
}

func (c *Circle) Area() {
	fmt.Println("Circle Area")
}

func (c *Circle) Perimeter() {
	fmt.Println("Circle Perimeter")
}

// 题目2
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	EmployeeID int
	person     Person
}

func (e *Employee) PrintInfo() {
	fmt.Println("Name", e.person.Name)
	fmt.Println("Age", e.person.Age)
	fmt.Println("EmployeeID", e.EmployeeID)
}

// Channel
// 题目1

func ChannelCaseInsert(chn chan int) {
	for i := 0; i < 10; i++ {
		chn <- i
	}
}
func ChannelCaseRead(chn chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(<-chn)
	}
}

// 题目2
func ChannelCache() {
	chn := make(chan int, 100)
	for i := 0; i < 100; i++ {
		chn <- i
	}

	go func(chn chan int) {
		for true {
			select {
			case a := <-chn:
				fmt.Println(a)
			default:
				fmt.Println("接收结束")
				return
			}
		}

	}(chn)
	time.Sleep(1 * time.Second)
}

// 锁机制
// 题目1
func TestChannel() {
	var k int
	mutex := sync.Mutex{}
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			for j := 0; j < 10000; j++ {
				k += 1
			}
			defer mutex.Unlock()
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(k)
}

// 题目2
func TestChannel2() {
	var k int64
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				atomic.AddInt64(&k, 1)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(k)
}

func main() {
	//a := 10
	//i := &a
	//pointInt(i)
	//fmt.Println(*i)

	//si := []int{1, 2, 3}
	//s := &si
	//pointSlice(s)
	//fmt.Println(*s)

	//go Go1to10()
	//go Go2to10()
	//time.Sleep(5 * time.Second)

	//goRoutineTask(func() {
	//	time.Sleep(1 * time.Second)
	//})
	//time.Sleep(3 * time.Second)

	//r := Rectangle{}
	//c := Circle{}
	//r.Area()
	//r.Perimeter()
	//c.Area()
	//c.Perimeter()
	//var sr Shape
	//sr = &r
	//sr.Area()
	//sr.Perimeter()
	//var sc Shape
	//sc = &c
	//sc.Area()
	//sc.Perimeter()

	//p := Person{
	//	"tree", 18,
	//}
	//e := Employee{person: p, EmployeeID: 1}
	//e.PrintInfo()

	//chn := make(chan int, 10)
	//go ChannelCaseInsert(chn)
	//time.Sleep(1 * time.Second)
	//go ChannelCaseRead(chn)
	//time.Sleep(1 * time.Second)

	//ChannelCache()

	//TestChannel()

	TestChannel2()

}
