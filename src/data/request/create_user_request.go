package request

type CreateUserRequest struct {
	Name string `validate:"required" json:"name"`
}
