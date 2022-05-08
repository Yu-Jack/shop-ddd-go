package usecase

import (
	"context"
)

type Saga interface {
	CheckoutSaga(ctx context.Context, groupId string, eventUUID string)
}
