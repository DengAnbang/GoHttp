package main

import (
	"main/gohttp"
	"main/goSocket"
	//_ "github.com/go-sql-driver/mysql"
)

func main() {
	host := "192.168.1.3"
	go goSocket.GoSocket(host, "9091")
	gohttp.GoHttp(host, "80")


}
