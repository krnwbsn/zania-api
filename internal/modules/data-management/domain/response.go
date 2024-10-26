package domain

import (
	"time"
	shareddomain "zania-api/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
)

// ResponseDataManagementList model
type ResponseDataManagementList struct {
	Meta candishared.Meta         `json:"meta"`
	Data []ResponseDataManagement `json:"data"`
}

// ResponseDataManagement model
type ResponseDataManagement struct {
	ID        int    `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Position  int    `json:"position"`
	ImageUrl  string `json:"imageUrl"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Serialize from db model
func (r *ResponseDataManagement) Serialize(source *shareddomain.DataManagement) {
	r.ID = source.ID
	r.Type = source.Type
	r.Title = source.Title
	r.Position = source.Position
	r.ImageUrl = source.ImageUrl
	r.CreatedAt = source.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = source.UpdatedAt.Format(time.RFC3339)
}

// ResponseUpdateSequenceDataManagement model
type ResponseUpdateSequenceDataManagement struct {
}
