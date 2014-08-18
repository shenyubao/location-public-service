package models

type JsonResult struct {
	Code    int
	Success bool
	Message string
	Data    interface {}
}

func Render(code int, message string, data interface {}) (result JsonResult) {
	_result := JsonResult{Code:code,Success:code == 100000, Message:message, Data:data}
	return _result
}

