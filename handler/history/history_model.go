package history

type CreateRequest struct {
	ProjectID      int    `json:"project_id,omitempty"`
	Description    string `json:"description,omitempty"`
	Url            string `json:"url,omitempty"`
	CurrentProcess string `json:"current_process,omitempty"`
}
