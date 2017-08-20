package auth

// Response is the standard response for auth
type Response struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}
