package main

import "net/http"

func handler(writer http.ResponseWriter,request *http.Request){
	print("haha")
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
	print("j")
}
