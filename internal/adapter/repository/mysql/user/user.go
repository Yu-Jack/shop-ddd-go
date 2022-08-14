package user

import (
	userUC "github.com/Yu-Jack/shop-ddd-go/internal/usecase/user"
	"gorm.io/gorm"
)

type repoUser struct {
	ID        string `json:"id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Amount    int    `json:"amount"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:milli"`
}

type repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) userUC.UserRepo {
	return &repo{
		db: db,
	}
}
