package util

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
通用返回结果
  - @param code
  - @param msg
  - @param data
  - @return
*/
func GeneralRes(code string, msg string, data interface{}) Result {
	r := Result{Code: code, Msg: msg, Data: data}
	return r
}

/*
成功返回结果
  - @param data
  - @return
*/
func SuccessRes(data interface{}) Result {
	r := Result{Code: "SUCCESS", Msg: "成功", Data: data}
	return r
}

/*
失败返回结果
  - @param msg
  - @param data
  - @return
*/
func FailRes(msg string, data interface{}) Result {
	r := Result{Code: "FAIL", Msg: msg, Data: data}
	return r
}
