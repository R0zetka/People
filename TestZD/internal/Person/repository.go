package Person

import (
	"context"
)

type Repositoriy interface {
	Create(ctx context.Context, person Person) error
	People(ctx context.Context) (m []Person, err error)
	OnePerson(ctx context.Context, id string) (Person, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, person Person) error
}
