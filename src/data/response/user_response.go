package response

import "github.com/satori/go.uuid"

type UserResponse struct {
	Id   uuid.UUID    `json:"id"`
	Name string `json:"name"`
}
