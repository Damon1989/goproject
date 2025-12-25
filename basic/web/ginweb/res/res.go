package res

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Success(data any) Response {
	return Response{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
}

func SuccessMsg(msg string) Response {
	return Response{
		Code: 0,
		Msg:  msg,
		Data: nil,
	}
}

func SuccessCodeMsg(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func Error(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}

func ErrorMsg(msg string) Response {
	return Response{
		Code: -1,
		Msg:  msg,
		Data: nil,
	}
}
func ErrorCode(code int) Response {
	return Response{
		Code: code,
		Msg:  "error",
		Data: nil,
	}
}

func ErrorCodeMsg(code int, msg string) Response {
	return Response{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
}
