package models

type ResponseBuilderValue struct {
	HttpStatusCode int    `json:"statusCode,omitempty"`
	Status         string `json:"status,omitempty"`
	Description    string `json:"description,omitempty"`
	ResponseCode   string `json:"responseCode,omitempty"`
	Messages       string `json:"messages,omitempty"`
}

type ResponseBuilderMessage struct {
	Language string `json:"language,omitempty"`
	Wording  string `json:"wording,omitempty"`
}

type ResponseBuilder struct {
	Regex string               `json:"regex,omitempty"`
	Value ResponseBuilderValue `json:"value,omitempty"`
}

func NewReponseBuilder() *ResponseBuilder {
	return &ResponseBuilder{}
}

type Responses struct {
	ResponseBuilder []*ResponseBuilder
}

func NewResponses() *Responses {
	return &Responses{}
}

func (r *Responses) AddResponseBuilder(rb *ResponseBuilder) {
	g := NewReponseBuilder()
	g.Regex = rb.Regex
	g.Value = rb.Value
	r.ResponseBuilder = append(r.ResponseBuilder, g)
}
