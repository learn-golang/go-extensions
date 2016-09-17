package main

// #cgo pkg-config: python3 --cflags --libs
// #define Py_LIMITED_API
// #include <Python.h>
import "C"
import "net/http"
import "fmt"
import "sync"
import "strings"


//export add
func add(left, right int) int {
	fmt.Println(left)
	fmt.Println(right)
	return left + right
}


//export goroutine
func goroutine(url string, lenth int) int {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println(url)
		url = strings.Trim(url, "%00%00%00")
		fmt.Println(url)
		resp, err := http.Get(url)
		fmt.Println(resp)
		fmt.Println(err)
		defer wg.Done()
	}()
	wg.Wait()
	return 0
}

func main() {
}
