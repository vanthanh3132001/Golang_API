package http1

type MyRequest struct {
	Id          string `json:"id,omitempty" db:"id, omitempty"`
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
	UpdatedTime string `json:"updated_time,omitempty" db:"updated_time, omitempty"`
	UpdatedBy   string `json:"updated_by,omitempty" db:"updated_by, omitempty"`
}
type HelloWorldRequest struct {
	Name string `json:"name"`
}
