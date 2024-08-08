package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
)

type LinkEndpoints struct {
	usecase usecases.LinkUsecase
}

func NewLinkEndpoints(usecase usecases.LinkUsecase) *LinkEndpoints {
	return &LinkEndpoints{usecase: usecase}
}

// GetUser godoc
// @Summary Get user by ID
// @Description Retrieve a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id query string true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user [get]
func (c *LinkEndpoints) GetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	http.Error(w, "Invalid user ID", http.StatusBadRequest)
	user, err := c.usecase.GetLinkByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// UserResponse represents the structure of the response for a user
type UserResponse struct {
	ID          string `json:"id"`
	IsDeleted   bool   `json:"is_deleted"`
	IsVisible   bool   `json:"is_visible"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
	CreatedAt   string `json:"created_at"`
	ModifiedAt  string `json:"modified_at"`
	Expiration  string `json:"expiration"`
	Clicks      int    `json:"clicks"`
}

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
	Message string `json:"message"`
}