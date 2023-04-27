package http1

import "time"

type MyResponse struct {
	Id          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	MetaData    string    `json:"metadata,omitempty"`
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	UpdatedBy   string    `json:"updated_by,omitempty"`
}

// HelloWorldResponse là một struct để lưu trữ thông tin response của API
type HelloWorldResponse struct {
	Message string `json:"message"`
}
