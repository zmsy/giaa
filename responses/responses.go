/*
Package responses is used to provide a standard format
for all JSON-web responses.
*/
package responses

import "encoding/json"

// Response is used to send back an error message
// in the event that something went wrong.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResponse is a function used to return a json.Marshal'd byte array rep
func SuccessResponse(data interface{}) ([]byte, error) {
	resp := Response{}
	resp.Status = "success"
	resp.Message = ""
	resp.Data = data

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return out, err
}

// ErrorResponse is used to pass back an empty container in the event that something
// could not be found.
func ErrorResponse(errorMessage string) ([]byte, error) {
	resp := Response{}
	resp.Status = "error"
	resp.Message = errorMessage
	resp.Data = nil

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return out, err
}
