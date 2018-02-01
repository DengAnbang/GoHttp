package main

import (

	//_ "github.com/go-sql-driver/mysql"
	"gohttp"
	"goSocket"
)

func main() {
	host := "localhost"
	//host := "192.168.1.3"
	go goSocket.GoSocket(host, "9091")
	gohttp.GoHttp(host, "80")

//fmt.Print("asdd")

}
