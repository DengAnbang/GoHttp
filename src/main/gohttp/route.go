package gohttp

import (
	"net/http"
	"encoding/json"
	"main/utils"
	"os"
)

func MainRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	file, i := os.Open("layout.html")
	defer file.Close()
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf)
	}
	utils.Loge(i)
}
func UserRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	utils.Loge(r)
	values := r.Form
	phone := values.Get("phone")
	password := values.Get("password")
	if phone == "123" && password == "123456" {
		data := Succeed(User{Name: " ", Age: 22})
		w.Write(data)
	} else {
		failure := Failure(1, "密码错误")
		w.Write(failure)
	}
	//for k, v := range values {
	//	utils.Loge(k)
	//	utils.Loge(v)
	//}

}

func Succeed(data interface{}) []byte {
	bytes, _ := json.Marshal(ResultData{Msg: "成功", Data: data})
	return bytes
}
func Failure(code int32, ssg string) []byte {
	bytes, _ := json.Marshal(ResultData{Msg: ssg, Code: code})
	return bytes
}
