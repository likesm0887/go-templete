package repository

import (
	memberData "awesomeProject1/member/adapter/repository/data"
)

type IMemberRepository interface {

	AddUserInformation(member memberData.Information) (*string, error)

	GetUserInfo(memberId string) (*memberData.Information, error)

	UpdateMemberInfo(data memberData.Information) error

	GetAll() (*[]memberData.Information, error)

}
