package gohttp

type ResultData struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type User struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
	Msg  string `json:"msg"`
}
