package model

// response default api
type Response struct {
	Data interface{} `json:"data,omitempty"`
	Err  error       `json:"error,omitempty"`
}
