package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

type LinkUsecase struct {
	linkRepo domain.LinkRepository
}

func NewLinkUseCase(linkRepo domain.LinkRepository) *LinkUsecase {
	return &LinkUsecase{linkRepo: linkRepo}
}
