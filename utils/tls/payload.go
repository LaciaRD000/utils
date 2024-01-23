package http

import (
	"bytes"
	"encoding/json"
)

func setPayload(body map[string]string) *bytes.Buffer {
	payload, _ := json.Marshal(body)
	return bytes.NewBuffer(payload)
}
