package main

import (
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "/Users/wuchengjin/Documents/yimi/src/learning/https/server.crt", "/Users/wuchengjin/Documents/yimi/src/learning/https/server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}