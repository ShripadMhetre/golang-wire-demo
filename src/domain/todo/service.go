package todo

import (
	"github.com/google/uuid"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/user"
)

type Service interface {
	GetRemoveRepo

	Add(createdBy user.User, text string) (*Todo, error)

	Set(id uuid.UUID, text string) (*Todo, error)
}
