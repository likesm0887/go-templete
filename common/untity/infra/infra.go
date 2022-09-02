package infra

import (
	"awesomeProject1/common/errors"
	"encoding/json"
	"github.com/xorcare/pointer"
	"log"
	"net/http"
	"time"
)

// WriteError
// 透過這個function來寫http error response，會記錄到Log
// @params writer http response 必填
// @params errorResponse 錯誤訊息 必填
// @params objects 要額外紀錄的物件 選填
func WriteError(writer http.ResponseWriter, errorResponse *errors.ErrorResponse, objects ...interface{}) {
	outPut := ""
	if len(objects) > 0 {
		for _, object := range objects {
			json, _ := json.Marshal(object)
			outPut += string(json)
		}
	}
	writer.WriteHeader(errorResponse.HttpStatusCode)
	writer.Write(errorResponse.ToJson())
	log.Printf("error : %v \n data:%s", errorResponse, outPut)
}

func WriteSuccess(writer http.ResponseWriter) {
	response := struct {
		Success        bool    `json:"success"`
		HttpStatusCode int     `json:"-"`                     // http的錯代碼
		ErrorCode      *string `json:"error_code" `           // 錯誤代碼
		StatusText     *string `json:"status_text,omitempty"` // 錯誤資訊(原生的錯誤訊息 Debug使用)
		Message        *string `json:"message,omitempty"`     // 錯誤訊息(使用者錯誤訊息)
	}{
		Success:        true,
		HttpStatusCode: http.StatusOK,
		ErrorCode:      nil,
		StatusText:     pointer.String("Success"),
	}
	writer.Write(ToJson(response))
	return
}

func WriteResponse(writer http.ResponseWriter, data []byte) bool {
	_, err := writer.Write(data)
	if err != nil {
		WriteError(writer, errors.NewInternalError(err))
		return false
	}
	return true
}

func ToJson(data interface{}) []byte {
	marshal, _ := json.Marshal(data)
	return marshal
}

func JsonTo(data string, result interface{}) {
	err := json.Unmarshal([]byte(data), result)
	if err != nil {
		return
	}
}

// GetTimeNow
// TimePattern "2006-01-02 15:04:05"
func GetTimeNow() string {
	timePattern := "2006-01-02 15:04:05"
	return time.Now().Format(timePattern)
}
