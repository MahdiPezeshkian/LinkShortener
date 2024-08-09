package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

func (u *LinkUsecase) SaveLink(link *domain.Link) error {
	return u.linkRepo.Insert(link)
}
