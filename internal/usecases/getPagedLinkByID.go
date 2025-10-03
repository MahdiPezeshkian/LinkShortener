package usecases

import (
	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (u *LinkUsecase) GetPagedLinkByID(pf *pkg.PaginationRequest) (data []*domain.LinkOutputDto, totalCount int, err error) {

	links, totalCount, err := u.linkRepo.GetPaged(pf)
	if err != nil {
		return nil, 0, err
	}

	for _, link := range links {
		dto := &domain.LinkOutputDto{
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
		data = append(data, dto)
	}
	return
}
