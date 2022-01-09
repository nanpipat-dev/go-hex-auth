package domain

// ===================== gorm model zone ===================== //
type Token struct {
	ID           int32  `gorm:"primary_key;auto_increment;not_null"`
	MemberID     string `gorm:"type:varchar(200);not null"`
	Member       Members
	RefreshToken string `gorm:"type:text;not null"`
	Expire       int32  `gorm:"type:integer;not null"`
}

func (t Token) TableName() string {
	return "token"
}

// ===================== gorm model zone ===================== //

// ===================== token model ===================== //

type TokenRequest struct {
	MemberID     string `json:"member_id"`
	RefreshToken string `json:"refresh_token"`
	Expire       int32  `json:"expire"`
}

// ===================== token model ===================== //
