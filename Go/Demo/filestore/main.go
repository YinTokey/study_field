package main

import (
	"filestore/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload",handler.UploadHander)
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	//http.HandleFunc("/file/query",handler.Que)
	http.HandleFunc("/file/download",handler.DownloadHandler)
	http.HandleFunc("/file/update",handler.UploadHander)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)


	http.ListenAndServe(":8080",nil)
}
