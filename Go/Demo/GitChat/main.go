package main


import (
	"net/http"
)

type HelloHandler struct {}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	print("xx")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	print("aaaa")
}


func main() {

	heHandler := HelloHandler{}
	woHandler := WorldHandler{}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
	}

	http.Handle("/hello",&heHandler)
	http.Handle("/world",&woHandler)

	server.ListenAndServe()
}
