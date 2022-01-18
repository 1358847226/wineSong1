package util

import "wineSong/entity"

func ErrReturn(code int, err error, content string) entity.Result {
	var res entity.Result
	res.Code = code
	res.Error = err
	res.Message = content
	return res
}
