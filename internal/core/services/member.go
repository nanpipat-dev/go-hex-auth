package services

import (
	"errors"
	"go-hex-auth/internal/core/domain"
	"go-hex-auth/internal/repositories"
	"go-hex-auth/package/security"
)

type MemberServiceInterface interface {
	CreateMember(m domain.MembersRequest) error
	GetMember(memberid string) (domain.MembersResponse, error)
	Login(username, password string) (domain.MembersResponse, error)
}

type MemberService struct {
	repo repositories.MemberRepositoryInterface
}

func NewMemberService(repo repositories.MemberRepositoryInterface) *MemberService {
	return &MemberService{
		repo: repo,
	}
}

func (s *MemberService) CreateMember(m domain.MembersRequest) error {
	uuid := security.GenerateUUID()
	pwd, err := security.EncryptPassword(m.Password)
	if err != nil {
		return err
	}

	member := domain.Members{
		MemberID:  uuid,
		Username:  m.Username,
		Password:  pwd,
		FirstName: m.FirstName,
		LastName:  m.LastName,
	}

	err = s.repo.CreateMember(member)
	if err != nil {
		return err
	}

	return nil
}

func (s *MemberService) GetMember(memberid string) (domain.MembersResponse, error) {
	member, err := s.repo.GetMember(memberid)
	if err != nil {
		return domain.MembersResponse{}, err
	}

	return member, nil
}

func (s *MemberService) Login(username, password string) (domain.MembersResponse, error) {
	check, err := s.repo.Login(username)
	if err != nil {
		return domain.MembersResponse{}, errors.New("invalid username")
	}

	err = security.VerifyPassword(check.Password, password)
	if err != nil {
		return domain.MembersResponse{}, errors.New("invalid password")
	}

	res := domain.MembersResponse{
		MemberID: check.MemberID,
	}

	return res, nil
}
