package controller

const (
	errorBadRequestBody  = "bad_request_body"
	errorBadRequestParam = "bad_request_param"
	errorBadRequest      = "bad_request"
)

type ResponseData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponseCodeMsg(code, msg string) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
	}
}
