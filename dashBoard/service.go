package dashBoard

import "context"

type MyService interface {
	GetByID(ctx context.Context, id string) (string, error)
	//Create(ctx context.Context, data string) error
	//Update(ctx context.Context, id string, data string) error
	//Delete(ctx context.Context, id string) error
}
