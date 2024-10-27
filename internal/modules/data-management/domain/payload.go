package domain

type DataManagementSequence struct {
	ID       int `json:"id"`
	Position int `json:"position"`
}

type UpdateDataManagementSequenceRequest struct {
	Sequence []DataManagementSequence `json:"sequence"`
}
