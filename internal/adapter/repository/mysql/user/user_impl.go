package user

import (
	"errors"

	userDomain "github.com/Yu-Jack/shop-ddd-go/internal/domain/user"
)

func (repo *repo) DecreaseUserAmount(userId string, orderAmount int) error {
	tx := repo.db.Begin()

	var c repoUser
	tx.Model(&repoUser{}).Where("id = ?", userId).Find(&c)

	if c.Amount < orderAmount {
		return errors.New("user doesn't have enough money")
	}
	c.Amount -= orderAmount
	tx.Save(c)

	result := tx.Commit()

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *repo) CreateUser(c *userDomain.User) error {
	user := repoUser{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Amount:    c.Amount,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
	result := repo.db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *repo) GetAllUsers() (cus []userDomain.User, err error) {
	var users []repoUser
	result := repo.db.Model(&repoUser{}).Find(&cus)

	if result.Error != nil {
		return cus, result.Error
	}

	for _, user := range users {
		cus = append(cus, userDomain.User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Amount:    user.Amount,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return cus, nil
}
