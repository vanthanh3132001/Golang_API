package postgresDB

type CreateUser struct {
	Name        string `json:"name,omitempty" db:"name, omitempty"`
	Description string `json:"description,omitempty" db:"description, omitempty"`
	MetaData    string `json:"metadata,omitempty" db:"metadata, omitempty"`
}
