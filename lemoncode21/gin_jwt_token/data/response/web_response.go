package response

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
}
