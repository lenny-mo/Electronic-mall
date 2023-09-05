package serializer

type Response struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   string      `json:"error"`
}

// TokenData
//
// 携带用户信息以及token的结构体
type TokenData struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}
