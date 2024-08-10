package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (c *LinkEndpoints) GetLink(ctx *gin.Context) {
	id := ctx.Query("id")

	linkDto, err := c.usecase.GetLinkByID(id)
	if err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusNotFound, "Not found")
		ctx.JSON(http.StatusNotFound, errResponse)
		return
	}

	output := pkg.NewRestApiResponse(&linkDto, http.StatusOK, "done")
	ctx.JSON(http.StatusOK, output)
}
