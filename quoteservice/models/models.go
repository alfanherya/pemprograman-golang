package models

import "time"

type AdditionalRequest struct {
	Uuid           string
	HttpStatusCode int
	Starttime      time.Time
}

func NewAdditionalRequest() *AdditionalRequest {
	return &AdditionalRequest{}
}

type AdditionalResponse struct {
	HttpStatusCode int
}

func NewAdditionalResponse() *AdditionalResponse {
	return &AdditionalResponse{}
}

type Param struct {
	ParamName  string `json:"paramName"`
	ParamValue string `json:"paramValue"`
}

func NewParam() *Param {
	return &Param{}
}
