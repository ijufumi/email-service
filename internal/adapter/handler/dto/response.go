package dto

// Result is response parameter
type Result struct {
	Status bool   `json:"status"`
	Error  *error `json:"error"`
}
