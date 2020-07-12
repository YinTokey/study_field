package main


import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func hello(writer http.ResponseWriter,request *http.Request, p httprouter.Params){
	//fmt.Fprint(writer,"参数为 %s",p.ByName("nae"))
	len := request.ContentLength
	body := make([]byte,len)
	request.Body.Read(body)
	fmt.Fprint(writer,string(body))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// 重定向
func headerExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}


func main() {

	//mux := httprouter.New()
	//mux.GET("/hello/:nae",hello)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
	}

	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)

	server.ListenAndServe()
}
