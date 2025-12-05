package workflows

type Node struct {
	ID     string                 `json:"id"`
	Type   string                 `json:"type"`
	Name   string                 `json:"name"`
	Params map[string]interface{} `json:"params"`
}

type Edge struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type WorkflowDefinition struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type CreateWorkflowInput struct {
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	DefinitionJSON WorkflowDefinition `json:"definition_json"`
}

type UpdateWorkflowInput struct {
	Name           *string             `json:"name"`
	Description    *string             `json:"description"`
	DefinitionJSON *WorkflowDefinition `json:"definition_json"`
}
type Workflow struct {
	ID             string             `json:"id"`
	OrgID          string             `json:"org_id"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	DefinitionJSON WorkflowDefinition `json:"definition_json"`
	CreatedAt      string             `json:"created_at"`
	UpdatedAt      string             `json:"updated_at"`
}
type ListWorkflowsParams struct {
	OrgID  string
	Limit  int32
	Offset int32
}
type ListWorkflowsResult struct {
	Workflows []Workflow `json:"workflows"`
	Total     int32      `json:"total"`
}
