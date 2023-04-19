package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

var Hostname = ""

func init() {
	Hostname, _ = os.Hostname()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "1")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"code\":1}")
	})

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		if Hostname == "" {
			fmt.Fprint(w, "无法获取主机名")
			return
		}

		fmt.Fprint(w, Hostname)
	})

	http.ListenAndServe(":80", nil)
}
