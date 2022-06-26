package dtos

type Response struct {
	Data interface{} `json:"data"`
	Meta struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"meta"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}
type AddRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Salary   int64  `json:"salary"`
}

type DeleteRequest struct {
}
type DeleteResponse struct {
}
type EditRequest struct {
}
type EditResponse struct {
}
