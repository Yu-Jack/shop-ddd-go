package entity

var db = []*Consumer{}

type repo struct{}

func New() *repo {
	return &repo{}
}

func (repo *repo) DecreaseConsumerAmount(consumerId string, orderAmount int) bool {
	for _, c := range db {
		if c.ID == consumerId && c.Amount > orderAmount {
			c.Amount = c.Amount - orderAmount
			return true
		}
	}
	return false
}

func (repo *repo) CreateConsumer(c Consumer) error {
	db = append(db, &c)
	return nil
}

func (repo *repo) GetAllConsumers() ([]*Consumer, error) {
	return db, nil
}
