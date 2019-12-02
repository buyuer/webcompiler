package webcompiler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func compile(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, _ := ioutil.ReadAll(req.Body)
		info := CompileInfo{}
		json.Unmarshal(body, &info)
		result := info.CompileAndRun()
		reqBody, _ := json.Marshal(&result)
		w.Write(reqBody)
	}
}

func favicon(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	file, err := ioutil.ReadFile("F:\\ProjectFiles\\JetBrains\\GoLand\\webcompiler\\static\\favicon.ico")
	fmt.Println(file)
	checkError(err)
	fmt.Println(len(file))
	w.Write(file)

}
