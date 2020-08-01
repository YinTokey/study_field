package handler

import (
	"encoding/json"
	"filestore/meta"
	"filestore/util"
	"fmt"
	io "io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHander(w http.ResponseWriter,r *http.Request) {
	if r.Method == "GET" {
		data,err := ioutil.ReadFile("./static/view/home.html")

		if err != nil {
			fmt.Println("server error")
			return
		}
		io.WriteString(w,string(data))

	} else if r.Method == "POST" {
		// 接收文件流,并转存到本地
		file,head,err := r.FormFile("file")
		if err != nil {
			fmt.Printf("get file error %s",err.Error())
			return;
		}
		defer file.Close()

		fileMeta := meta.FileMeta {
			FileName : head.Filename,
			Location : "/tmp/" +head.Filename,
			UploadAt : time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile,err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("fail to create file, error :  %s",err.Error())
			return
		}

		defer newFile.Close()

		 io.Copy(newFile,file)

		//if err != nil {
		//	fmt.Printf("fail to save, error :  %s",err.Error())
		//	return
		//}

		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)


	}
}

//  文件信息查询
func UploadSucHandler(w http.ResponseWriter,r *http.Request) {
	io.WriteString(w,"upload finish")
}

func GetFileMetaHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	filehash := r.Form["filehash"][0]
	fMeta := meta.GetFileMeta(filehash)
	data,err := json.Marshal(fMeta)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter,r *http.Request) {

	r.ParseForm()
	fsha1 := r.Form.Get("filehash")

	fm := meta.GetFileMeta(fsha1)

	f,err := os.Open(fm.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer f.Close()

	data,err := ioutil.ReadAll(f)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("Content-Description","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)
}
