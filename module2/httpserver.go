package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func healthzFunc(w http.ResponseWriter, r *http.Request) {
	code := "200"
	w.Write([]byte(code))
}
func httpAccessFunc(w http.ResponseWriter, r *http.Request) {
	if len(r.Header) > 0 {
		for k, v := range r.Header {
			w.Header().Set(k, v[0])
		}
	}
	os.Setenv("VERSION", "V1.0")

	//2. 获取环境变量"VERSION"
	name := os.Getenv("VERSION")
	fmt.Println("VERSION Env: ", name)

	//打印ip
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("获取ip错误:err,", err)
	}
	//判断ip
	if net.ParseIP(ip) != nil {
		fmt.Println("ip地址为:\n", ip)
	}
	fmt.Println("http 状态码:\n", http.StatusOK)
	log.Println(http.StatusOK)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "httpAccessFunc")
}
func main() {
	http.HandleFunc("/", httpAccessFunc)
	http.HandleFunc("/healthzFunc", healthzFunc)
	http.ListenAndServe(":3000", nil)
}
