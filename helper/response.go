package helper

import "strings"

// RESPONSE IS USED FOR STATIC SHAPE JSON RETURN
type Response struct {
	Status		bool 		`json:"status"`
	Message		string		`json:"message"`
	Errors		interface{}	`json:"errors"`
	Data 		interface{}	`json:"data"`
}

// EMPTY OBJECT IS USED WHEN DATA DOESNT WANT TO BE NULL ON JSON
type EmptyObj struct{}

// BUILD RESPONSE METHOD IS TO INJECT DATA VALUE TO DYNAMIC SUCCESS RESPONSE
func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status	: status,
		Message	: message,
		Errors	: nil,
		Data	: data,
	}
	return res
}

// BUILD ERROR RESPONSE METHOD IS TO INJECT DATA VALUE TO DYNAMIC FAILED RESPONSE
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status	: false,
		Message	: message,
		Errors	: splittedError,
		Data	: data,
	}
	return res
}