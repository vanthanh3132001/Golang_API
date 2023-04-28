package http1

type MyResponse struct {
	Id          string `json:"id,omitempty" db:"id, omitempty"`
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
	UpdatedTime string `json:"updated_time,omitempty" db:"updated_time, omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty" db:"updated_by, omitempty"`
}

// HelloWorldResponse là một struct để lưu trữ thông tin response của API
type HelloWorldResponse struct {
	Message string `json:"message"`
}
