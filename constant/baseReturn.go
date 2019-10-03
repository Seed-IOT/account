package constant

import "encoding/json"

// BaseReturn 1
type BaseReturn struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
