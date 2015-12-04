package main

import (
	"html/template"
	"net"
	"net/http"
)

type IP struct {
	Name    string
	Address string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tmpl/visitor.html")
	name := "nobody"
	if len(r.URL.Path[1:]) > 0 {
		name = r.URL.Path[1:]
	}
	ip := IP{Name: name, Address: getLocalIP()}
	tmpl.Execute(w, ip)
}

func johnnyHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("tmpl/johnny.html")
	ip := IP{Address: getLocalIP()}
	tmpl.Execute(w, ip)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	http.HandleFunc("/johnny", johnnyHandler)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
