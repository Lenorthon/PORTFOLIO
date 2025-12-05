package workflows

import "errors"

func ValidateWorkflow(def WorkflowDefinition) error {
	nodeIDs := map[string]bool{}

	for _, node := range def.Nodes {
		if node.ID == "" {
			return errors.New("node id missing")
		}
		if node.Type == "" {
			return errors.New("node type missing")
		}
		nodeIDs[node.ID] = true
	}

	for _, e := range def.Edges {
		if !nodeIDs[e.Source] {
			return errors.New("edge source does not exist: " + e.Source)
		}
		if !nodeIDs[e.Target] {
			return errors.New("edge target does not exist: " + e.Target)
		}
	}

	return nil
}
