package services

import (
	"errors"
	"fmt"
	"go-hex-auth/internal/core/domain"
	"go-hex-auth/internal/repositories"
	"go-hex-auth/package/security"
	"time"
)

const SecretKey = "secret"

type MemberServiceInterface interface {
	CreateMember(m domain.MembersRequest) error
	GetMember(memberid string) (domain.MembersResponse, error)
	Login(username, password string) (domain.LoginResponse, error)
	CheckRefreshToken(refresh string) (bool, error)
	Refresh(refresh string) (domain.LoginResponse, error)
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

func (s *MemberService) Login(username, password string) (domain.LoginResponse, error) {
	check, err := s.repo.Login(username)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid username")
	}

	err = security.VerifyPassword(check.Password, password)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid password")
	}

	token, err := security.NewToken(check.MemberID)

	if err != nil {
		return domain.LoginResponse{}, err
	}

	refresh, err := security.RefreshToken(check.MemberID)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	expire := time.Now().Add((time.Minute * 60)).Unix()

	refreshObj := domain.Token{
		MemberID:     check.MemberID,
		RefreshToken: refresh,
		Expire:       int32(expire),
	}

	err = s.repo.CreateToken(refreshObj)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	// fmt.Println(token, "token")
	// resclaims, err := security.ParseToken(token)
	// if err != nil {
	// 	return domain.LoginResponse{}, err
	// }

	// fmt.Println(resclaims.Id, "resclaims")

	res := domain.LoginResponse{
		AccessToken:  token,
		RefreshToken: refresh,
	}

	return res, nil
}

func (s *MemberService) CheckRefreshToken(refresh string) (bool, error) {
	refreshTk, err := security.RefreshCheck(refresh)
	if err != nil {
		return false, err
	}
	return s.repo.GetRefreshToken(refresh, refreshTk)
}

func (s *MemberService) Refresh(refresh string) (domain.LoginResponse, error) {
	refreshTk, err := security.RefreshCheck(refresh)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid token")
	}

	fmt.Println(refreshTk, "refresh")

	check, err := s.repo.GetRefreshToken(refresh, refreshTk)
	if err != nil {
		return domain.LoginResponse{}, errors.New("invalid token")
	}

	if !check {
		return domain.LoginResponse{}, errors.New("seesion expired")
	}

	token, err := security.NewToken(refreshTk)

	if err != nil {
		return domain.LoginResponse{}, err
	}

	refresh_token, err := security.RefreshToken(refreshTk)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	expire := time.Now().Add(time.Minute * 1).Unix()

	refreshObj := domain.Token{
		MemberID:     refreshTk,
		RefreshToken: refresh_token,
		Expire:       int32(expire),
	}

	err = s.repo.CreateToken(refreshObj)
	if err != nil {
		return domain.LoginResponse{}, err
	}
	res := domain.LoginResponse{
		AccessToken:  token,
		RefreshToken: refresh,
	}

	return res, nil

}
