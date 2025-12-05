package workflows

import (
	"context"
	"encoding/json"

	"github.com/Lenorthon/PORTFOLIO/backend/internal/db"
)

type Service struct {
	q *db.Queries
}

func NewService(q *db.Queries) *Service {
	return &Service{q}
}

func (s *Service) Create(ctx context.Context, orgID string, input CreateWorkflowInput) (db.Workflow, error) {

	if err := ValidateWorkflow(input.DefinitionJSON); err != nil {
		return db.Workflow{}, err
	}

	defBytes, _ := json.Marshal(input.DefinitionJSON)

	arg := db.CreateWorkflowParams{
		OrgID:          orgID,
		Name:           input.Name,
		Description:    input.Description,
		DefinitionJson: defBytes,
	}

	wf, err := s.q.CreateWorkflow(ctx, arg)
	return wf, err
}

func (s *Service) Get(ctx context.Context, id string) (db.Workflow, error) {
	return s.q.GetWorkflow(ctx, id)
}

func (s *Service) List(ctx context.Context, orgID string) ([]db.Workflow, error) {
	return s.q.ListWorkflows(ctx, orgID)
}

func (s *Service) Update(ctx context.Context, id string, input UpdateWorkflowInput) (db.Workflow, error) {
	wf, err := s.q.GetWorkflow(ctx, id)
	if err != nil {
		return wf, err
	}

	if input.DefinitionJSON != nil {
		if err := ValidateWorkflow(*input.DefinitionJSON); err != nil {
			return wf, err
		}
		defBytes, _ := json.Marshal(input.DefinitionJSON)
		wf.DefinitionJson = defBytes
	}

	if input.Name != nil {
		wf.Name = *input.Name
	}

	if input.Description != nil {
		wf.Description = *input.Description
	}

	updated, err := s.q.UpdateWorkflow(ctx, db.UpdateWorkflowParams{
		ID:             wf.ID,
		Name:           wf.Name,
		Description:    wf.Description,
		DefinitionJson: wf.DefinitionJson,
	})

	return updated, err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.q.DeleteWorkflow(ctx, id)
}
