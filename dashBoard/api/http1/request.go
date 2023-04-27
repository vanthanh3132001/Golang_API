package http1

import "time"

type MyRequest struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	MetaData    string    `json:"metadata,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
}
type HelloWorldRequest struct {
	Name string `json:"name"`
}
