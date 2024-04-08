package request

import "github.com/satori/go.uuid"

type UpdateUserRequest struct {
	Id   uuid.UUID    `validate:"required" json:"id"`
	Name string `validate:"required, min:3, max:250" json:"name"`
}
