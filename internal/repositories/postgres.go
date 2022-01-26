package repositories

import (
	"errors"
	"fmt"
	"go-hex-auth/internal/core/domain"
	"time"

	"gorm.io/gorm"
)

type MemberRepositoryInterface interface {
	CreateMember(m domain.Members) error
	GetMember(memberid string) (domain.MembersResponse, error)
	Login(username string) (domain.Members, error)
	CreateToken(token domain.Token) error
	GetRefreshToken(refresh, memberid string) (bool, error)
}

type MemberRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *MemberRepository {
	domain.MigrateDatabase(db)
	return &MemberRepository{
		db: db,
	}
}

func (r *MemberRepository) CreateMember(m domain.Members) error {
	tx := r.db.Create(&m)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *MemberRepository) CreateToken(token domain.Token) error {
	tx := r.db.Create(&token)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *MemberRepository) GetMember(memberid string) (domain.MembersResponse, error) {
	member := domain.Members{}
	condition, err := whereCondition()
	if err != nil {
		fmt.Println("err")
	}
	fmt.Println(condition)
	tx := r.db.Where("member_id = ?", memberid).First(&member)
	// tx := r.db.Where(condition).Find(&member)
	if tx.Error != nil {
		return domain.MembersResponse{}, tx.Error
	}

	res := domain.MembersResponse{
		MemberID:  member.MemberID,
		Username:  member.Username,
		FirstName: member.FirstName,
		LastName:  member.LastName,
	}

	return res, nil
}

func whereCondition() (map[string]interface{}, error) {
	expressions := map[string]interface{}{}
	id := "0a97eed9-88df-481d-aae6-f20df9ebe24c"
	name := "toptest"
	last := ""

	if id != "" {
		expressions["member_id like"] = "%" + id + "%"
	}
	if name != "" {
		expressions["username"] = name
	}

	if last != "" {
		expressions["last"] = last
	}

	return expressions, nil
}

func (r *MemberRepository) Login(username string) (domain.Members, error) {
	member := domain.Members{}
	tx := r.db.Where(&domain.Members{Username: username}).First(&member)
	if tx.Error != nil {
		return domain.Members{}, tx.Error
	}
	return member, nil
}

func (r *MemberRepository) GetRefreshToken(refresh, memberid string) (bool, error) {
	token := domain.Token{}
	now := time.Now().Unix()
	rmv := r.db.Where("expire < ?", now).Delete(&token)
	if rmv.Error != nil {
		return false, rmv.Error
	}
	tx := r.db.Where(&domain.Token{MemberID: memberid, RefreshToken: refresh}).First(&token)
	if tx.Error != nil {
		return false, tx.Error
	}

	if (int64(token.Expire) - time.Now().Unix()) <= 0 {
		return false, errors.New("expire")
	}

	return true, nil
}
