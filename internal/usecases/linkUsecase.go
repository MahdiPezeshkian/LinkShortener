package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

type LinkUsecase struct {
	linkRepo domain.LinkRepository
}

func NewLinkUseCase(linkRepo domain.LinkRepository) *LinkUsecase {
	return &LinkUsecase{linkRepo: linkRepo}
}

func (u *LinkUsecase) GetLinkByID(id string) (*domain.Link, error) {
	return u.linkRepo.FindByID(id)
}

func (u *LinkUsecase) SaveLink(link *domain.Link) error {
	return u.linkRepo.Insert(link)
}
