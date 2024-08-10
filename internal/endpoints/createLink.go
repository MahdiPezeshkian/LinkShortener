package endpoints

import (
	"net/http"

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

	dto, err := c.usecase.SaveLink(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.SetRestApiError[domain.LinkOutputDto](http.StatusConflict, err.Error()))
		return
	}

	output := pkg.NewRestApiResponse(&dto, http.StatusCreated, "done")

	ctx.JSON(http.StatusCreated, output)
}
