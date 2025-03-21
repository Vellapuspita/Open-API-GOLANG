package main 
import (
	"fmt"
	"net/http"
)

fun handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World!")
}

func main(){
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil)
}