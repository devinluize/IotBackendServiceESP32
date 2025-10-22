package responses

type ApiResponseError struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Success    bool        `json:"success"`
	Err        interface{} `json:"data"`
}
