package main

import (
	"fmt"
	"io"
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "httpAccessFunc")
}

//获取环境变量"VERSION"
func getVersion(w http.ResponseWriter, r *http.Request) {
	os.Setenv("VERSION", "V1.0")
	name := os.Getenv("VERSION")
	fmt.Println("VERSION Env: ", name)
	io.WriteString(w, name)
}

//打印ip
func getIP(w http.ResponseWriter, r *http.Request) {

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println("获取ip错误:err,", err)
	}
	if net.ParseIP(ip) != nil {
		fmt.Println("ip地址为:", ip)
	}
	io.WriteString(w, ip)
}
func getCode(w http.ResponseWriter, r *http.Request) {
	code := "200"
	fmt.Println("code:", code)
	w.Write([]byte(code))
}
func main() {
	http.HandleFunc("/", httpAccessFunc)
	http.HandleFunc("/healthzFunc", healthzFunc)
	http.HandleFunc("/version", getVersion)
	http.HandleFunc("/ip", getIP)
	http.HandleFunc("/code", getCode)
	fmt.Printf("http server is listening")
	http.ListenAndServe(":80", nil)
}
