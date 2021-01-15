package response

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Write(w http.ResponseWriter, data interface{}, status int) http.ResponseWriter {
	w.WriteHeader(status)
	res := Response{
		Code:    status,
		Message: http.StatusText(status),
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)

	return w
}

// 自定义的与业务相关的Code码
const (
	CodeClientError = 4001
	CodeServerError = 5001
)

var codeText = map[int]string{
	CodeClientError: "自定义客户端响应",
	CodeServerError: "自定义服务端响应",
}

func CodeText(code int) string {
	return codeText[code]
}

func WriteWithCode(w http.ResponseWriter, data interface{}, code int, message string) http.ResponseWriter {
	if message == "" {
		message = CodeText(code)
	}
	status, _ := strconv.Atoi(strconv.Itoa(code)[0:3])
	w.WriteHeader(status)
	res := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
	return w
}
