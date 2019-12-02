package webcompiler

import "net/http"

func Run() {
	file := http.FileServer(http.Dir("F:\\ProjectFiles\\JetBrains\\GoLand\\webcompiler\\static"))
	http.Handle("/", http.StripPrefix("/", file))
	http.Handle("/compile", http.HandlerFunc(compile))
	//http.Handle("/favicon.ico",http.HandlerFunc(favicon))
	err := http.ListenAndServe("localhost:80", nil)
	checkError(err)
}
