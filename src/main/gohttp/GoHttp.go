package gohttp

import (
	"net/http"
	"log"
	"main/utils"
)

func GoHttp(host, port string)  {
	utils.Loge("GoHttp")
	http.HandleFunc("/",MainRoute)
	http.HandleFunc("/app",UserRoute)
	err := http.ListenAndServe(host+":"+port, nil)
	if err!=nil {
		log.Fatal("err:",err)
	}
}
