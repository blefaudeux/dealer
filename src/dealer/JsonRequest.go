package dealer

import (
	"encoding/json"
)

type JsonRequest struct {
	Id   string           `json:"Id"`
	Data *json.RawMessage `json:"data"`
}
