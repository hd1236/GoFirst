package main

import (
	"HelloWorld/FileListingServer/FileListing"
	"log"
	"net/http"
	"os"
)

type appHandler func(writer http.ResponseWriter,
	request *http.Request) error

func errWrapper(handler appHandler) func(w http.ResponseWriter,
	r *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error:%s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code),
				code)

		}
	}
}
func main() {
	http.HandleFunc("/list/",
		errWrapper(FileListing.HandlerFileList))
	er := http.ListenAndServe(":8888", nil)
	if er != nil {
		panic(er)
	}
}
