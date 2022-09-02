package member_router

import (
	"awesomeProject1/common/errors"
	"awesomeProject1/common/untity/infra"
	"awesomeProject1/member/adapter/contract"
	"awesomeProject1/member/domain/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

type MemberRouter struct {
	memberService service.IMemberService
}

func NewMemberRouter(memberService service.IMemberService) *MemberRouter {
	return &MemberRouter{
		memberService: memberService,
	}
}

func (memberRouter *MemberRouter) GetApi(router *mux.Router) {
	memberRouter.setToPublicApi(router)
}

func (memberRouter *MemberRouter) setToPublicApi(router *mux.Router) {
	router.Path("/api/v1/information/{id}").HandlerFunc(memberRouter.getUserInfo).Methods(http.MethodGet)
	router.Path("/api/v1/information").HandlerFunc(memberRouter.AddUserInfo).Methods(http.MethodPost)
	router.Path("/api/v1/information").HandlerFunc(memberRouter.updateUserInfo).Methods(http.MethodPut)
}

// @Summary 更新使用者資料
// @Tags Member
// @Accept application/json
// @version 1.0
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param userId body login_jwt.User true "使用者資料"
// @Success 200 {string} string "ok"
// @Failure 404 {object} errors.ErrorResponse "1004 此帳號尚未註冊"
// @Failure 500 {object} errors.ErrorResponse "9999 內部錯誤"
// @Router /member/information/{id} [post]
func (memberRouter *MemberRouter) updateUserInfo(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	updateUserInfo := &contract.Information{}
	err := json.Unmarshal(reqBody, updateUserInfo)

	if err != nil {
		infra.WriteError(writer, errors.NewInternalError(err))
		return
	}

	updateUserInfoErr := memberRouter.memberService.UpdateUserInfo(*updateUserInfo)
	if updateUserInfoErr != nil {
		infra.WriteError(writer, updateUserInfoErr)
		return
	} else {
		infra.WriteSuccess(writer)
	}
}

// @Summary 取得使用者資料
// @Tags Member
// @Accept application/json
// @version 1.0
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Param userId query string true "登入資料"
// @Success 200 {object} login_jwt.User "User info"
// @Failure 404 {object} errors.ErrorResponse "1004 此帳號尚未註冊"
// @Failure 500 {object} errors.ErrorResponse "9999 內部錯誤"
// @Router /member [get]
func (memberRouter *MemberRouter) getUserInfo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id := params["id"]
	user, err := memberRouter.memberService.GetUserInfo(id)
	if err != nil {
		infra.WriteError(writer, err, id)
	} else {
		if infra.WriteResponse(writer, infra.ToJson(user)) {
			return
		}
	}
}

func (memberRouter *MemberRouter) AddUserInfo(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	updateUserInfo := &contract.Information{}
	err := json.Unmarshal(reqBody, updateUserInfo)
	if err != nil {
		infra.WriteError(writer, errors.NewInternalError(err))
		return
	}

	addUserInfoErr := memberRouter.memberService.AddUserInfo(*updateUserInfo)
	if addUserInfoErr != nil {
		infra.WriteError(writer, addUserInfoErr)
	} else {
		infra.WriteSuccess(writer)

	}
}

func (memberRouter *MemberRouter) UpdateUserInfo(writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	updateUserInfo := &contract.Information{}
	err := json.Unmarshal(reqBody, updateUserInfo)
	if err != nil {
		infra.WriteError(writer, errors.NewInternalError(err))
		return
	}

	updateUserInfoErr := memberRouter.memberService.UpdateUserInfo(*updateUserInfo)
	if err != nil {
		infra.WriteError(writer, updateUserInfoErr)
	} else {
		infra.WriteSuccess(writer)
	}
}
