package model

// response default api
type Response struct {
	Data interface{} `json:"data,omitempty" swaggerignore:"true"`
	Err  error       `json:"error,omitempty" swaggerignore:"true"`
}
