package contextor

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

const (
	ERROR        = 500
	SUCCESS      = 200
	UNAUTHORIZED = 401
)
