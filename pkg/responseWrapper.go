package pkg

type ResponseBase struct {
	Message      string `json:"message,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	HasError     bool   `json:"has_error"`
}

type RestApiResponse[T any] struct {
	ResponseBase
	Value      *T  `json:"value,omitempty"`
	StatusCode int `json:"status_code"`
}

func NewRestApiResponse[T any](value *T, statusCode int, message string) *RestApiResponse[T] {
	return &RestApiResponse[T]{
		ResponseBase: ResponseBase{
			Message:      message,
			ErrorMessage: "empty",
			HasError:     false,
		},
		Value:      value,
		StatusCode: statusCode,
	}
}

func SetRestApiError[T any](statusCode int, errorMessage string) *RestApiResponse[T] {
	return &RestApiResponse[T]{
		ResponseBase: ResponseBase{
			Message:      "empty",
			ErrorMessage: errorMessage,
			HasError:     true,
		},
		Value:      nil,
		StatusCode: statusCode,
	}
}

type PagedRestApiResponse[T any] struct {
	ResponseBase
	Items       []T  `json:"items,omitempty"`
	PageNumber  int  `json:"page_number"`
	PageSize    int  `json:"page_size"`
	ItemsCount  int  `json:"items_count"`
	TotalCount  int  `json:"total_count"`
	HasNextPage bool `json:"has_next_page"`
	StatusCode  int  `json:"status_code"`
}

func NewPagedRestApiResponse[T any](items []T, pageNumber, pageSize, totalCount, statusCode int, message string) *PagedRestApiResponse[T] {
	itemsCount := len(items)
	hasNextPage := (pageNumber*pageSize)+itemsCount < totalCount

	return &PagedRestApiResponse[T]{
		ResponseBase: ResponseBase{
			Message:      message,
			ErrorMessage: "empty",
			HasError:     false,
		},
		Items:       items,
		PageNumber:  pageNumber,
		PageSize:    pageSize,
		ItemsCount:  itemsCount,
		TotalCount:  totalCount,
		HasNextPage: hasNextPage,
		StatusCode:  statusCode,
	}
}

func SetPagedRestApiError[T any](statusCode int, errorMessage string) *PagedRestApiResponse[T] {
	return &PagedRestApiResponse[T]{
		ResponseBase: ResponseBase{
			Message:      "empty",
			ErrorMessage: errorMessage,
			HasError:     true,
		},
		Items:       nil,
		PageNumber:  0,
		PageSize:    0,
		ItemsCount:  0,
		TotalCount:  0,
		HasNextPage: false,
		StatusCode:  statusCode,
	}
}
