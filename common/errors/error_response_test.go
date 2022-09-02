package errors

import (

	"fmt"
	"testing"
)

func TestCommonError_NewDataBaseError(t *testing.T) {
	fmt.Println(string(LoginError.UserExistResp.ToJson()))
	fmt.Println(string(LoginError.UserRegisterFailResp.ToJson()))
}


