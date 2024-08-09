package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

func (u *LinkUsecase) GetLinkByID(id string) (*domain.Link, error) {
	return u.linkRepo.FindByID(id)
}
