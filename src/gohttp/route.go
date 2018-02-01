package gohttp

import (
	"net/http"
	"encoding/json"

	"utils"
	"html/template"
	"fmt"
)

func MainRoute(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("src/main/templates/layout.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	t.Execute(w, nil)
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
