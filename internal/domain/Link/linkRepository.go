package domain

type LinkRepository interface {
	Insert(user *Link) error
	FindByID(id string) (*Link, error)
	FindAll() ([]*Link, error)
	Delete(id string) error
	Update(link *Link) error
}
