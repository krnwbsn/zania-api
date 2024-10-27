package domain

type DataManagementSequence struct {
	ID       int `json:"id"`
	Position int `json:"sequence"`
}

type UpdateDataManagementSequenceRequest struct {
	Sequence []DataManagementSequence `json:"sequence"`
}
