package responses

type ErrorResponses struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Err        error       `json:"-"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
}

type StandarAPIResponses struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Err        error       `json:"-"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data"`
}
