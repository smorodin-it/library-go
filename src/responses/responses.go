package responses

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

type ResponseAdd struct {
	Id string `json:"id"`
}

type ResponseStatus struct {
	Status bool `json:"status"`
}
