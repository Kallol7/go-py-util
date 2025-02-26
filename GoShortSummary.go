package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func findMax(ara []int64) int64 {
	var max int64 = 0

	// range gives (index, value) pairs
	for _, v := range ara {
		if v > max {
			max = v
		}
	}
	return max
}

type person struct {
	name string
	age  int
}

func count(s string) {
	for i := 0; i < 2; i++ {
		fmt.Print(s, i, "\n")
		time.Sleep(time.Nanosecond)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		sum := 0
		for i := 0; i < n; i++ {
			sum += i
		}
		results <- sum
	}
}

func concurrence() {
	sendChan := make(chan int)
	recvChan := make(chan int)

	go worker(sendChan, recvChan)

	for i := 0; i < 5; i++ {
		sendChan <- i
		fmt.Println(<-recvChan)
	}
}

// generic
func double[T int | float64 | string](inp T) string {
	var s string
	if fmt.Sprintf("%T", inp) == "int" {
		s = fmt.Sprint("Integer type ")
	} else {
		s = fmt.Sprint("Float type ")
	}
	s += fmt.Sprint(inp + inp)
	return s
}

func main() {
	// go count("abc")
	// go count("def")
	// // to prevent main goroutine from stopping its execution
	// // goroutine will continue to run until you press enter
	// fmt.Scanln()

	// c := make(chan string)
	// go func() {
	// 	c <- "abc"
	// }()
	// dff := <-c
	// fmt.Println(dff)

	// runConcurrence := true
	// if runConcurrence {
	// 	concurrence()
	// 	return
	// }

	g1 := 32
	g2 := 32.57
	fmt.Println(double(g1)) // Output: Integer type 64
	fmt.Println(double(g2)) // Output: Float type 65.14

	var a float32 = 3.4218374242345435
	var b float64 = 6.57
	x := a + float32(b)
	fmt.Println(x)

	fmt.Println(findMax([]int64{3, 5, 999, 45}))

	// pointer
	blf := new(person) // returns pointer to person
	func(ptr *person) {
		ptr.name = "new_person"
		(*ptr).age = 35 // *ptr.age will give error
	}(blf)
	fmt.Println(blf) // Output: &{new_person 35}

	// append only works on slice, i.e., array without length
	pson := []person{{name: "ABC", age: 3}}
	pson = append(pson, person{age: 99})
	fmt.Println(pson)
	fmt.Println(pson[0])

	l := []int{3, 45, 6, 4, 5}
	l = append(l, 3)
	fmt.Println(l)

	d := make(map[string]int)
	d["one"] = 1
	d["two"] = 2
	d["three"] = 3
	delete(d, "three")

	for key, val := range d {
		fmt.Printf("Key: %v, Value: %v\n", key, val)
	}

	// for i := 0; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	var arr [5]string
	arr[0] = "df"
	arr[2] = "dfs23"
	fmt.Println((len((arr))))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Input was: %q\n", line)
	}
}
