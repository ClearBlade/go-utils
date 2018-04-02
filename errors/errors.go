package errors

import (
	"encoding/json"
	"fmt"

	cc "github.com/clearblade/go-utils/errors/constants/category"
	uuid "github.com/clearblade/go-utils/uuid"
)

type Info struct {
	Id            uuid.Uuid           `json:"id"`
	Code          float64             `json:"code"`
	Level         float64             `json:"level"`
	Category      cc.CategoryConstant `json:"category"`
	Message       string              `json:"message"`
	Detail        string              `json:"detail"`
	LowLevelError error               `json:"lowLevelError"`
	Line          string              `json:"line"`
}

type Response struct {
	Info       Info    `json:"error" mapstructure:"error"`
	StatusCode float64 `json:"statusCode"`
}

func (e *Response) Error() string {
	marsha, _ := json.Marshal(e)
	return string(marsha)
}

func (e *Response) HTTPStatusCode() float64 {
	return e.StatusCode
}

func (e *Response) FromNet(body interface{}) *Response {
	var imAMap map[string]interface{}
	var ok bool
	if imAMap, ok = body.(map[string]interface{}); !ok {
		return createEmptyErrorResponse(body)
	}

	var errorInfo map[string]interface{}

	if errorInfo, ok = imAMap["error"].(map[string]interface{}); !ok {
		return createEmptyErrorResponse(body)
	}
	var statusCode float64
	if statusCode, ok = imAMap["statusCode"].(float64); !ok {
		return createEmptyErrorResponse(body)
	}

	// alright, it must be the format we're expecting. let's grab the fields we need

	parsedId, err := uuid.ParseUuid(errorInfo["id"].(string))
	if err != nil {
		return createEmptyErrorResponse(body)
	}
	info := Info{
		Id:            parsedId,
		Message:       errorInfo["message"].(string),
		Code:          errorInfo["code"].(float64),
		Level:         errorInfo["level"].(float64),
		Category:      cc.CategoryConstant(errorInfo["category"].(string)),
		Detail:        errorInfo["detail"].(string),
		LowLevelError: fmt.Errorf("%+v", errorInfo["lowLevelError"]),
		Line:          errorInfo["line"].(string),
	}

	return &Response{
		Info:       info,
		StatusCode: statusCode,
	}
}

func createEmptyErrorResponse(body interface{}) *Response {
	info := Info{
		Message: fmt.Sprintf("%+v", body),
	}
	return &Response{
		Info:       info,
		StatusCode: 0,
	}
}

func CreateResponseFromNet(body interface{}) *Response {
	var errResp Response
	return errResp.FromNet(body)
}
