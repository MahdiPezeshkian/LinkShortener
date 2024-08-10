package endpoints

import (
	"net/http"
	"time"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
	"github.com/gin-gonic/gin"
)

func (c *LinkEndpoints) CreateLink(ctx *gin.Context) {
	var input domain.LinkInputDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.SetRestApiError[domain.LinkOutputDto](http.StatusBadRequest, "Invalid input"))
		return
	}

	if input.OriginalURL == "" {
		ctx.JSON(http.StatusBadRequest, pkg.SetRestApiError[domain.LinkOutputDto](http.StatusBadRequest, "Original URL is required"))
		return
	}

	link := domain.NewLink(input.OriginalURL, time.Now().AddDate(0, 1, 0))

	if err := c.usecase.SaveLink(link); err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.SetRestApiError[domain.LinkOutputDto](http.StatusInternalServerError, "Failed to save link"))
		return
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

	output := pkg.NewRestApiResponse(&dto, http.StatusCreated, "done")

	ctx.JSON(http.StatusCreated, output)
}
