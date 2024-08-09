package domain

import "github.com/MahdiPezeshkian/LinkShortener/pkg"

type LinkRepository interface {
	Insert(user *Link) error
	FindByID(id string) (*Link, error)
	FindOneByCondition(condition string, args ...interface{}) (*Link, error)
	FindAll() ([]*Link, error)
	Delete(id string) error
	Update(link *Link) error
	HardDelete(id string) error
	GetPaged(pagination *pkg.PaginationRequest) ([]*Link, int, error)
}
