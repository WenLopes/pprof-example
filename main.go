package main

import (
	"net/http"
	"strconv"

	_ "net/http/pprof"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Handle("/cpu", http.HandlerFunc(CPUIntensive))
	go http.ListenAndServe(":8080", r)
	http.ListenAndServe(":6060", nil)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func CPUIntensive(w http.ResponseWriter, r *http.Request) {
	n := 42
	result := fib(n)
	w.Write([]byte(strconv.Itoa(result)))
}

func MemoryIntensive(w http.ResponseWriter, r *http.Request) {
	num := 1000000000
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		slice[i] = i
	}
	w.Write([]byte(strconv.Itoa(num)))
}
