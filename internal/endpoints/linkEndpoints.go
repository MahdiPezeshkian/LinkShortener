package endpoints

import (
	"github.com/MahdiPezeshkian/LinkShortener/internal/usecases"
)

type LinkEndpoints struct {
	usecase usecases.LinkUsecase
}

func NewLinkEndpoints(usecase usecases.LinkUsecase) *LinkEndpoints {
	return &LinkEndpoints{usecase: usecase}
}
