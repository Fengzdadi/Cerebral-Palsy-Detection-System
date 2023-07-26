package Serializer

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

type Datalist struct {
	Item  interface{} `json:"item"`
	Total int         `json:"total"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}
