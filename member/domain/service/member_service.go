package service

import (
	"awesomeProject1/common/errors"
	"awesomeProject1/member/adapter/contract"
	"awesomeProject1/member/domain/mapper"
	"awesomeProject1/member/domain/repository"
)

type IMemberService interface {


	GetUserInfo(userId string) (*contract.Information, *errors.ErrorResponse)

	UpdateUserInfo( info contract.Information) *errors.ErrorResponse

	AddUserInfo(info contract.Information) *errors.ErrorResponse
}

type MemberService struct {
	memberRepository repository.IMemberRepository
}

func NewMemberService( memberRepository repository.IMemberRepository) *MemberService {
	service := &MemberService{
		memberRepository: memberRepository,
	}

	return service
}

func (memberService *MemberService) GetUserInfo(userId string) (*contract.Information, *errors.ErrorResponse) {
	userInformation, err := memberService.memberRepository.GetUserInfo(userId)

	if err != nil {
		return nil, errors.LoginError.UserNotFoundCodeResp
	}
	userContract := mapper.ToContractInfo(*userInformation)
	return &userContract, nil
}


func (memberService *MemberService) UpdateUserInfo( info contract.Information) *errors.ErrorResponse {

	err := memberService.memberRepository.UpdateMemberInfo(mapper.ToDataInfo(info))
	if err != nil {
		return errors.NewInternalError(err)
	}
	return nil
}

func (memberService *MemberService) AddUserInfo( info contract.Information) *errors.ErrorResponse {

	_,err := memberService.memberRepository.AddUserInformation(mapper.ToDataInfo(info))
	if err != nil {
		return errors.NewInternalError(err)
	}
	return nil
}