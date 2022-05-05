package entity

type Consumer struct {
	ID        string `json:"id" gorm:"primarykey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Amount    int    `json:"amount"`
	CreatedAt int    `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int    `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
