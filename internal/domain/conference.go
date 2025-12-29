package domain

// Conference represents a conference where talks are submitted and presented.
type Conference struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}
