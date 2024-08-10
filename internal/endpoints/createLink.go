package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
)

func (c *LinkEndpoints) CreateLink(w http.ResponseWriter, r *http.Request) {
	var input domain.LinkInputDto
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusMethodNotAllowed, "Method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusBadRequest, "Invalid input")
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	if input.OriginalURL == "" {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusBadRequest, "Original URL and Short URL are required")
		w.WriteHeader(http.StatusBadRequest)

		if err := json.NewEncoder(w).Encode(errResponse); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	link := domain.NewLink(input.OriginalURL,  time.Now().AddDate(0, 1, 0))

	if err := c.usecase.SaveLink(link); err != nil {
		errResponse := pkg.SetRestApiError[domain.LinkOutputDto](http.StatusInternalServerError, "Failed to save link")
		w.WriteHeader(http.StatusInternalServerError)

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

	output := pkg.NewRestApiResponse(&dto, http.StatusCreated, "done")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
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
