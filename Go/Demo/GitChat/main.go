package main

import (
	"net/http"
	"html/template"
)


func handler(writer http.ResponseWriter,request *http.Request){
	print("haha")
}

func index(writer http.ResponseWriter,request *http.Request){
	files := []string{"templates/layout.html","templates/navbar.html","templates/index.html"}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads();if err == nil {
		templates.ExecuteTemplate(w,"layout",threads)
	}
	print("response")
}

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/",http.StripPrefix("/static/",files))
	mux.HandleFunc("/",index)

	server := &http.Server{
		Addr: "0.0.0.0:8000",
		Handler:mux,
	}


	server.ListenAndServe()
}
