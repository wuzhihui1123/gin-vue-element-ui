package web

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS         = 200
	CLIENT_ERROR    = 400
	PERMISSION_DENY = 403
	SERVER_ERROR    = 500
)

var MsgFlags = map[int]string{
	SUCCESS:      "success",
	SERVER_ERROR: "服务端错误",
	CLIENT_ERROR: "客户端请求错误",
}

func Success(data interface{}) (ret Response) {
	ret = Response{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	}
	return
}
func ServerError(msg string) (ret Response) {
	ret = Response{
		Code: SERVER_ERROR,
		Msg:  GetMsg(SERVER_ERROR),
	}
	return
}
func ClientError(msg string) (ret Response) {
	ret = Response{
		Code: CLIENT_ERROR,
		Msg:  GetMsg(CLIENT_ERROR),
	}
	return
}
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[SERVER_ERROR]
}
