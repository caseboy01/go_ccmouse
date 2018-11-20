package main

import (
	"log"
	"mooc_google_shizhan/errhandling/filelistingserver/filelisting"
	"net/http"
	"os"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {

			//log.Warn("Error handling request %s",err.Error())  //github gopm

			log.Printf("Error occurred"+"handling request: %s", err.Error()) //内置的

			code := http.StatusOK //200
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound //404
			case os.IsPermission(err):
				code = http.StatusForbidden //403
			default:
				code = http.StatusInternalServerError //500
			}

			http.Error(w, http.StatusText(code), code)
		}
	}

}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
