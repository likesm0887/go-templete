package errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

// ErrorResponse example
type ErrorResponse struct {
	HttpStatusCode int    `json:"-"`                     // http的錯代碼
	ErrorCode      string `json:"error_code" `           // 錯誤代碼
	StatusText     string `json:"status_text,omitempty"` // 錯誤資訊(原生的錯誤訊息 Debug使用)
	Message        string `json:"message,omitempty"`     // 錯誤訊息(使用者錯誤訊息)
	Success        bool   `json:"success"`
}

func NewError(text string) error {
	return errors.New(text)
}

// 會員相關 1

// region Login 登入相關錯誤 錯誤代碼 1

// LoginError 登入相關
var LoginError = struct {
	LoginFailResp        *ErrorResponse
	UserExistResp        *ErrorResponse
	UserRegisterFailResp *ErrorResponse
	UserNotFoundCodeResp *ErrorResponse
}{
	LoginFailResp:        &ErrorResponse{Success:false,HttpStatusCode: http.StatusUnauthorized, ErrorCode: loginFailCode, StatusText: "登入失敗", Message: "登入失敗"},
	UserExistResp:        &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: accountExistCode, StatusText: "此帳號已存在", Message: "此帳號已存在"},
	UserRegisterFailResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusInternalServerError, ErrorCode: registerFailCode, StatusText: "註冊失敗", Message: "註冊失敗"},
	UserNotFoundCodeResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusNotFound, ErrorCode: accountNotFoundCode, StatusText: "此帳號尚未註冊", Message: "此帳號尚未註冊"},
}

const (
	// LoginFailCode 登入失敗
	loginFailCode = "1001"
	// MemberExistCode 此帳號已存在
	accountExistCode = "1002"
	// MemberRegisterFailCode 註冊失敗
	registerFailCode = "1003"
	// MemberNotFoundCode 此帳號尚未註冊
	accountNotFoundCode = "1004"
)

// endregion

//  region VideoCall 視訊通話相關錯誤 錯誤代碼5

var VideoCallError = struct {
	RequestAccessTokenResp *ErrorResponse
	CreateRoomResp         *ErrorResponse
}{
	RequestAccessTokenResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: RequestAccessTokenError, StatusText: "要求AccessToken錯誤", Message: "要求AccessToken錯誤"},
	CreateRoomResp:         &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: CreateRoomError, StatusText: "建立房間錯誤", Message: "建立房間錯誤"},
}

const (
	// RequestAccessTokenError 要求AccessToken錯誤
	RequestAccessTokenError = "5001"
	// CreateRoomError 建立房間錯誤
	CreateRoomError = "5002"
)

// endregion

// region Member 會員相關錯誤 錯誤代碼 6

var MemberError = struct {
	AddFavoriteCounselorRepeatResp    *ErrorResponse
	PlaylistNotFoundResp              *ErrorResponse
	MusicNotFoundResp                 *ErrorResponse
	AddMusicInPlaylistMusicRepeatResp *ErrorResponse
	ResetPasswordResp *ErrorResponse
}{
	AddFavoriteCounselorRepeatResp:    &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: AddFavoriteCounselorRepeat, StatusText: "重複加入", Message: "重複加入"},
	PlaylistNotFoundResp:              &ErrorResponse{Success:false,HttpStatusCode: http.StatusNotFound, ErrorCode: PlaylistNotFoundResp, StatusText: "找不到播放清單", Message: "找不到播放清單"},
	AddMusicInPlaylistMusicRepeatResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusConflict, ErrorCode: AddMusicInPlaylistMusicRepeat, StatusText: "此音樂已經加入過了", Message: "此音樂已經加入過了"},
	MusicNotFoundResp:                 &ErrorResponse{Success:false,HttpStatusCode: http.StatusNotFound, ErrorCode: MusicNotFound, StatusText: "找不到音樂", Message: "找不到音樂"},
	ResetPasswordResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusConflict, ErrorCode: ResetPasswordFail, StatusText: "驗證失敗", Message: "驗證失敗"},
}

const (
	// AddFavoriteCounselorRepeat 重複加入
	AddFavoriteCounselorRepeat = "6001"
	// PlaylistNotFoundResp 找不到Playlist
	PlaylistNotFoundResp = "6002"
	// AddMusicInPlaylistMusicRepeat 找不到Playlist
	AddMusicInPlaylistMusicRepeat = "6003"
	// MusicNotFound 找不到音樂
	MusicNotFound = "6004"
	// ResetPasswordFail 忘記密碼驗證碼錯誤
	ResetPasswordFail = "6005"
)

// endregion

// region Appointment 預約相關錯誤 錯誤代碼 7

var AppointmentError = struct {
	AppointmentNotFoundResp        *ErrorResponse
	AppointmentAlreadyAcceptedResp *ErrorResponse
	CancellationTimeExceededResp   *ErrorResponse
	VideoRoomNotYetEstablishedResp *ErrorResponse
}{
	AppointmentNotFoundResp:        &ErrorResponse{Success:false,HttpStatusCode: http.StatusNotFound, ErrorCode: AppointmentNotFound, StatusText: "查無此預約", Message: "查無此預約"},
	AppointmentAlreadyAcceptedResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: AppointmentAlreadyAccepted, StatusText: "此預約已經被接受", Message: "此預約已經被接受"},
	CancellationTimeExceededResp:   &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: CancellationTimeExceeded, StatusText: "超過可取消時間", Message: "超過可取消時間"},
	VideoRoomNotYetEstablishedResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: VideoRoomNotYetEstablished, StatusText: "諮商房間尚未建立", Message: "諮商房間尚未建立"},
}

const (
	// AppointmentNotFound 查無此預約
	AppointmentNotFound = "7001"
	// AppointmentAlreadyAccepted 此預約已接受
	AppointmentAlreadyAccepted = "7002"
	// AppointmentDataNotComplete 預約內容不完整
	AppointmentDataNotComplete = "7003"
	// CancellationTimeExceeded 超過可取消時間
	CancellationTimeExceeded = "7004"
	// VideoRoomNotYetEstablished 諮商房間尚未建立
	VideoRoomNotYetEstablished = "7005"
)

// endregion

// region Counselor 諮商師相關錯誤 錯誤代碼 8

var CounselorError = struct {
	AddCounselorInInstitutionRepeatResp *ErrorResponse
}{
	AddCounselorInInstitutionRepeatResp: &ErrorResponse{Success:false,HttpStatusCode: http.StatusBadRequest, ErrorCode: AddCounselorInInstitutionRepeat, StatusText: "重複加入", Message: "重複加入"},
}

const (
	// AddCounselorInInstitutionRepeat 重複加入諮商師
	AddCounselorInInstitutionRepeat = "8001"
)

// endregion

// region Common 常見錯誤 error 錯誤代碼 9

const (
	// InternalError 內部錯誤
	internalError = "9999"
	// HttpError Http錯誤
	httpError = "9998"
	// DataBaseError 資料庫錯誤
	dataBaseError = "9997"
	// paramsError 參數錯誤
	paramsError = "9996"
	// RequestTooLargeError 檔案太大
	RequestTooLargeError = "9994"
	// UploadError 檔案上傳失敗
	UploadError = "9993"
	// DownloadError 檔案下載失敗
	DownloadError = "9992"
	// RequestBodyError 請求內容錯誤
	RequestBodyError = "9991"
)

func NewDownloadError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      DownloadError,
		StatusText:     error.Error(),
		Message:        "檔案下載失敗",
	}
}

func NewAppointmentDataNotCompleteError(field string) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusBadRequest,
		ErrorCode:      AppointmentDataNotComplete,
		StatusText:     "內容不完整: " + field,
		Message:        field + "錯誤",
	}
}

func NewRequestBodyError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      RequestBodyError,
		StatusText:     "請求內容錯誤 請檢查api文件",
		Message:        "請求內容錯誤",
	}
}

func NewUploadError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      UploadError,
		StatusText:     error.Error(),
		Message:        "上傳檔案失敗",
	}
}

func NewUploadLargeError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusRequestEntityTooLarge,
		ErrorCode:      RequestTooLargeError,
		StatusText:     error.Error(),
		Message:        "上傳檔案太大",
	}
}

func NewParamsError(params string, errors ...error) *ErrorResponse {

	if len(errors) != 0 {
		return &ErrorResponse{
			Success:false,
			HttpStatusCode: http.StatusBadRequest,
			ErrorCode:      paramsError,
			StatusText:     errors[0].Error(),
			Message:        params + " 參數錯誤",
		}
	} else {
		return &ErrorResponse{
			Success:false,
			HttpStatusCode: http.StatusBadRequest,
			ErrorCode:      paramsError,
			StatusText:     params + " 參數錯誤",
			Message:        params + " 參數錯誤",
		}
	}

}

func NewInternalError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      internalError,
		StatusText:     error.Error(),
		Message:        "發生內部錯誤",
	}
}

func NewHttpError(error http.Response) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      httpError,
		StatusText:     error.Status,
		Message:        "http發生錯誤，status code:" + strconv.Itoa(error.StatusCode),
	}
}

func NewDataBaseError(error error) *ErrorResponse {
	return &ErrorResponse{
		Success:false,
		HttpStatusCode: http.StatusInternalServerError,
		ErrorCode:      dataBaseError,
		StatusText:     error.Error(),
		Message:        "資料庫發生錯誤",
	}
}

// endregion

func (errorResponse *ErrorResponse) ToJson() []byte {
	json, _ := json.Marshal(errorResponse)
	return json
}
