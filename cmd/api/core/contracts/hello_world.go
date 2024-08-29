package contracts

type HelloWorldRequest struct {
	UserID   string `json:"x-user-id" header:"x-user-id" validate:"required"`
	DateFrom string `json:"date_from" form:"date_from" validate:"required"`
	DateTo   string `json:"date_to" form:"date_to" validate:"required"`
}
