package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (c *LinkEndpoints) GetPagedLinks(w http.ResponseWriter, r *http.Request) {

	pageNumberStr := r.URL.Query().Get("page_number")
	pageSizeStr := r.URL.Query().Get("page_size")
	sortOrder := r.URL.Query().Get("sort_order")

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

	if r.Method != http.MethodGet {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusMethodNotAllowed, "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	linksDto, totalCount, err := c.usecase.GetPagedLinkByID(paginationRequest)
	if err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusInternalServerError, "Failed to fetch links")
		w.WriteHeader(http.StatusInternalServerError)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	output := pkg.NewPagedRestApiResponse(linksDto, paginationRequest.PageNumber, paginationRequest.PageSize, totalCount, http.StatusOK, "done")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusInternalServerError, "Failed to encode response")
		w.WriteHeader(http.StatusInternalServerError)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}
}
