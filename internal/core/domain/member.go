package domain

import (
	"gorm.io/gorm"
)

// ===================== gorm model zone ===================== //
func MigrateProfileImage(db *gorm.DB) {
	if db == nil {
		panic("An error when connect database")
	}

	db.AutoMigrate(Members{})
}

type Members struct {
	MemberID  string `gorm:"primaryKey;type:varchar(200)"`
	Username  string `gorm:"type:varchar(200)"`
	Password  string `gorm:"type:varchar(200)"`
	FirstName string `gorm:"type:varchar(200)"`
	LastName  string `gorm:"type:varchar(200)"`
}

func (t Members) TableName() string {
	return "members"
}

// ===================== gorm model zone ===================== //

// ===================== member model ===================== //

type MembersRequest struct {
	MemberID  string `json:"member_id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type MembersResponse struct {
	MemberID  string `json:"member_id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// ===================== member model ===================== //
