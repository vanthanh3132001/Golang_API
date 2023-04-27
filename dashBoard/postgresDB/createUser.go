package postgresDB

type CreateUser struct {
	Name        string `json:"fullName,omitempty" db:"name, omitempty"`
	Description string `json:"password,omitempty" db:"description, omitempty"`
	MetaData    string `json:"email,omitempty" db:"metadata, omitempty"`
}
