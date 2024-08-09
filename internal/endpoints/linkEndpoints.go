package endpoints

import (
	"encoding/json"
	"net/http"
	"time"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
)

type LinkEndpoints struct {
	usecase usecases.LinkUsecase
}

func NewLinkEndpoints(usecase usecases.LinkUsecase) *LinkEndpoints {
	return &LinkEndpoints{usecase: usecase}
}

func (c *LinkEndpoints) CreateLink(w http.ResponseWriter, r *http.Request) {
	var input domain.LinkOutputDto

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// خواندن و اعتبارسنجی ورودی
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// بررسی اینکه URLهای ورودی معتبر باشند
	if input.OriginalURL == "" || input.ShortURL == "" {
		http.Error(w, "Original URL and Short URL are required", http.StatusBadRequest)
		return
	}

	// ایجاد یک انتیتی جدید از مدل Link
	link := domain.NewLink(input.OriginalURL, input.ShortURL, time.Now().AddDate(0, 1, 0)) // فرض بر این است که Expiration یک ماه بعد است.

	// ذخیره لینک در دیتابیس یا هر ذخیره‌سازی مورد نیاز
	if err := c.usecase.SaveLink(link); err != nil {
		http.Error(w, "Failed to save link", http.StatusInternalServerError)
		return
	}

	// آماده‌سازی و برگرداندن خروجی
	output := domain.LinkOutputDto{
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (c *LinkEndpoints) GetLink(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	link, err := c.usecase.GetLinkByID(id)
	if err != nil {
		http.Error(w, "Link not found", http.StatusNotFound)
		return
	}

	output := domain.LinkOutputDto{
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
	json.NewEncoder(w).Encode(output)
}
