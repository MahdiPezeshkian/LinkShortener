package endpoints

import (
	"net/http"
	"strconv"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
	"github.com/gin-gonic/gin"
)

func (c *LinkEndpoints) GetPagedLinks(ctx *gin.Context) {

	pageNumberStr := ctx.Param("page_number")
	pageSizeStr := ctx.Param("page_size")
	sortOrder := ctx.Param("sort_order")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		pageNumber = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	if sortOrder == "" {
		sortOrder = "asc"
	}

	paginationRequest := &pkg.PaginationRequest{
		PageNumber: pageNumber,
		PageSize:   pageSize,
		SortOrder:  sortOrder,
	}

	linksDto, totalCount, err := c.usecase.GetPagedLinkByID(paginationRequest)
	if err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusInternalServerError, "Failed to fetch links")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	output := pkg.NewPagedRestApiResponse(linksDto, paginationRequest.PageNumber, paginationRequest.PageSize, totalCount, http.StatusOK, "done")
	ctx.JSON(http.StatusOK, output)
}
