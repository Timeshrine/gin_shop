package serializer

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type Token struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type Datalist struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: Datalist{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}