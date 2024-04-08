package response

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
