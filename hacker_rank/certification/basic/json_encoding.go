package basic

import (
	"bytes"
	"encoding/json"
	"io"
)

type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

var marshal = json.Marshal

func EncodeManager(manager *Manager) (io.Reader, error) {
	b, err := marshal(manager)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
