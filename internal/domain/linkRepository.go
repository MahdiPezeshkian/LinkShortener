package domain

import (
	"github.com/MahdiPezeshkian/LinkShortener/internal/domain/entity"
)

type LinkRepository interface {
	Insert(user *entity.Link) error
	FindByID(id string) (*entity.Link, error)
	FindAll() ([]*entity.Link, error)
	Delete(id string) error
	Update(link *entity.Link) error
}
