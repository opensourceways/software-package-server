package controller

const (
	errorBadRequest       = "bad_request"
	errorBadRequestBody   = "bad_request_body"
	errorBadRequestParam  = "bad_request_param"
	errorBadRequestHeader = "bad_request_header"
	errorBadRequestCookie = "bad_request_cookie"
)

type ResponseData struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponseData(data interface{}) ResponseData {
	return ResponseData{
		Data: data,
	}
}

func newResponseCodeError(code string, err error) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  err.Error(),
	}
}

func newResponseCodeMsg(code, msg string) ResponseData {
	return ResponseData{
		Code: code,
		Msg:  msg,
	}
}

func NewBadRequestHeader(msg string) ResponseData {
	return ResponseData{
		Code: errorBadRequestHeader,
		Msg:  msg,
	}
}

func NewBadRequestCookie(msg string) ResponseData {
	return ResponseData{
		Code: errorBadRequestCookie,
		Msg:  msg,
	}
}

func NewBadRequest(msg string) ResponseData {
	return ResponseData{
		Code: errorBadRequest,
		Msg:  msg,
	}
}
