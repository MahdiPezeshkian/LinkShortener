package usecases

import (
	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (u *LinkUsecase) GetPagedLinkByID(pf *pkg.PaginationRequest) ([]*domain.LinkOutputDto, int, error) {

	links, tCount, err := u.linkRepo.GetPaged(pf)

	var dtos []*domain.LinkOutputDto
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
		dtos = append(dtos, dto)
	}
	return dtos, tCount, err
}
