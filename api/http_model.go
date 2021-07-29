package api

type ResponseDTO struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *ResponseDTO) Success(msg string) *ResponseDTO {
	r.Code = 200
	r.Message = msg
	return r
}

func (r *ResponseDTO) SuccessWithData(msg string, data interface{}) *ResponseDTO {
	r.Code = 200
	r.Message = msg
	r.Data = data
	return r
}
