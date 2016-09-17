package main

// #cgo pkg-config: python3 --cflags --libs
// #define Py_LIMITED_API
// #include <Python.h>
import "C"
import "net/http"
import "fmt"
import "sync"


//export add
func add(left, right int) int {
	fmt.Println(left)
	fmt.Println(right)
	return left + right
}


//export goroutine
func goroutine(url string, length int) int {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println(len(url))
		fmt.Println(length)
		fmt.Println(url)
		resp, err := http.Get(url)
		fmt.Println(resp)
		fmt.Println(err)
		defer wg.Done()
	}()
	wg.Wait()
	return 0
}

//export python_getter
func python_getter(url string, length int) int {
	fmt.Println(url)
	resp, err := http.Get(url)
	fmt.Println(resp)
	fmt.Println(err)
	return 0
}

func main() {
}
