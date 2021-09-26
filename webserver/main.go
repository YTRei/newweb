package main 

import (
	"io"
	"net/http"
)

func print1to10()int{
	sum := 0
	for i :=1;i<=10;i++{
		sum +=i
	}
	return sum
}
func firstPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello, this is my fisrt page!!!</h1>")
}

func main() {
	http.HandleFunc("/", firstPage)
	http.ListenAndServe(":8000", nil)
}