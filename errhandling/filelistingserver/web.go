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

		//recover
		defer func() {

			if r := recover(); r != nil {
				log.Printf("Panic:%v", r)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(w, r)
		if err != nil {

			//log.Warn("Error handling request %s",err.Error())  //github gopm

			log.Printf("Error occurred"+"handling request: %s", err.Error()) //内置的

			if userErr, ok := err.(userError); ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}

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

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
