package exception

import "encoding/json"

// ExFields Created custom errors that are subscribed to the Errors interface
type ExFields map[string]interface{}

func (e ExFields) Error() string {
	value, _ := json.Marshal(e)
	return string(value)
}
