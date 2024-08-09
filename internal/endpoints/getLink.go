package endpoints

import (
	"encoding/json"
	"net/http"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (c *LinkEndpoints) GetLink(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method != http.MethodGet {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusMethodNotAllowed, "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	link, err := c.usecase.GetLinkByID(id)
	if err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusNotFound, "Not found")
		w.WriteHeader(http.StatusNotFound)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
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
	output := pkg.NewRestApiResponse(&dto, 200, "done")

	json.NewEncoder(w).Encode(output)
}
