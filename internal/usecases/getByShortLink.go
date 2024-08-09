package usecases

import domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"

func (u *LinkUsecase) GetByShortLink(sh string) (*domain.LinkOutputDto, error) {

	link, err := u.linkRepo.FindOneByCondition("short_url = ?", sh)

	if err != nil {
		return nil, err
	}

	link.Click()
	u.linkRepo.Update(link)

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
