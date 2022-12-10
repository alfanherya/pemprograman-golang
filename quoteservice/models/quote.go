package models

import (
	"time"
)

type QuoteRq struct {
	Id           int       `json:"id,omitempty"`
	Nama         string    `json:"nama,omitempty"`
	Quote        string    `json:"quote,omitempty"`
	ModifiedBy   string    `json:"modifiedby,omitempty"`
	CreatedDate  time.Time `json:"createddate" db:"SYS_CREATION_DATE"`
	ModifiedDate time.Time `json:"modifieddate" db:"SYS_CREATION_DATE"`
}

type QuoteRs struct {
	ErrorCode    string            `json:"code"`
	StatusCode   string            `json:"status"`
	ErrorMessage string            `json:"message"`
	Data         DataRs            `json:"data"`
	Internal     AdditionalRequest `json:"-"`
}

type DataRs struct {
	Nama       string `json:"nama,omitempty"`
	Quote      string `json:"quote,omitempty"`
	ModifiedBy string `json:"modified,omitempty"`
}
