package nodit

import (
	"fmt"
)

// Error represents an API error
type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("nodit: status %d: %s", e.StatusCode, e.Message)
} 