package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// 判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func cHandler(w http.ResponseWriter, r *http.Request) {

	// 默认数据
	fileName := "data/api" + strings.Replace(r.URL.Path, "/", "_", -1) + ".json"
	jsonDat := `{
	"errno":200,
	"msg":"succ",
	"data":""
}`

	if checkFileIsExist(fileName) {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			jsonDat = "error"
		}
		jsonDat = string(data)
	} else {
		d1 := []byte(jsonDat)
		ioutil.WriteFile(fileName, d1, 0644)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	io.WriteString(w, jsonDat)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cHandler)

	http.ListenAndServe(":12345", mux)
}
