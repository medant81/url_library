package author

import "time"

type Author struct {
	Id        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
}
