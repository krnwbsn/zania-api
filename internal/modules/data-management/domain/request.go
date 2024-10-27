package domain

import (
	shareddomain "zania-api/pkg/shared/domain"
)

// RequestDataManagement model
type RequestDataManagement struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Title    string `json:"title"`
	Position int    `json:"position"`
}

// Deserialize to db model
func (r *RequestDataManagement) Deserialize() (res shareddomain.DataManagement) {
	res.Type = r.Type
	res.Title = r.Title
	res.Position = r.Position
	return
}

// RequestUpdateSequenceDataManagement model
type RequestUpdateSequenceDataManagement struct {
}
