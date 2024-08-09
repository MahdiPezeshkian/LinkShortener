package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

func (u *LinkUsecase) GetLinkByID(id string) (*domain.LinkOutputDto, error) {

	link, err := u.linkRepo.FindByID(id)

	dto := domain.LinkOutputDto{
		Id:          link.Id,
		Isdeleted:   link.Isdeleted,
		IsVisibled:  link.IsVisibled,
		OriginalURL: link.OriginalURL,
		ShortURL:    link.ShortURL,
		CreatedAt:   link.CreatedAt,
		ModifiedAt:  link.ModifiedAt,
		Expiration:  link.Expiration,
		Clicks:      link.Clicks,
	}
	return &dto, err
}
