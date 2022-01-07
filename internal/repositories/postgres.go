package repositories

import (
	"go-hex-auth/internal/core/domain"

	"gorm.io/gorm"
)

type MemberRepositoryInterface interface {
	CreateMember(m domain.Members) error
	GetMember(memberid string) (domain.MembersResponse, error)
	Login(username string) (domain.Members, error)
}

type MemberRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *MemberRepository {
	domain.MigrateProfileImage(db)
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

func (r *MemberRepository) GetMember(memberid string) (domain.MembersResponse, error) {
	member := domain.Members{}
	tx := r.db.Where("member_id = ?", memberid).First(&member)
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

func (r *MemberRepository) Login(username string) (domain.Members, error) {
	member := domain.Members{}
	tx := r.db.Where(&domain.Members{Username: username}).First(&member)
	if tx.Error != nil {
		return domain.Members{}, tx.Error
	}
	return member, nil
}
