package usecases

import (
	"errors"
	"time"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
)

func (u *LinkUsecase) SaveLink(input *domain.LinkInputDto) (domain.LinkOutputDto, error) {
	link := domain.NewLink(input.OriginalURL, time.Now().AddDate(0, 1, 0))

	exist, err := u.linkRepo.FindManyByCondition("short_url = ? or original_url = ?", link.ShortURL, link.OriginalURL)

	if err != nil {
		return domain.LinkOutputDto{}, err
	}

	if len(exist) > 0 {
		for _, li := range exist {
			if li.ShortURL == link.ShortURL {
				link = domain.NewLink(input.OriginalURL, time.Now().AddDate(0, 1, 0))
			} else if li.OriginalURL == link.OriginalURL {
				err = errors.New("This Link already exists")
			}
		}

		if err != nil {
			return domain.LinkOutputDto{}, err
		}
	}

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

	err = u.linkRepo.Insert(link)

	if err != nil {
		return domain.LinkOutputDto{}, err
	}

	return dto, nil
}
