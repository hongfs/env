package main

import (
	"fmt"
	"io"
	"log"
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
		fmt.Fprint(w, "1\n")
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "{\"code\":1}\n")
	})

	http.HandleFunc("/output", func(w http.ResponseWriter, r *http.Request) {
		log.Println("接收到一个请求")

		for k, v := range r.Header {
			log.Println(k, v)
		}

		fmt.Fprint(w, "{\"code\":1}\n")
	})

	http.HandleFunc("/echo", func(writer http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)

		if err != nil {
			fmt.Fprint(writer, "读取请求体失败")
			return
		}

		log.Printf("接收到一个请求: %s\n", string(body))

		fmt.Fprint(writer, string(body))
	})

	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		if Hostname == "" {
			fmt.Fprint(w, "无法获取主机名")
			return
		}

		fmt.Fprint(w, Hostname+"\n")
	})

	http.ListenAndServe(":80", nil)
}
