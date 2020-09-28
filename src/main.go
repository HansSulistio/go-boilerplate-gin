package main

import(
	"net/http"
	"fmt"
	"github.com/hans/go-boilerplate-gin/src/controller/index"
)

func main(){
	http.HandleFunc("/hello", helloserver)	
	fmt.Println("Starting Server on Port :8080")
	http.ListenAndServe(":8080",nil)
}
func helloserver(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Hello World!" + index.Hai)
		
}

