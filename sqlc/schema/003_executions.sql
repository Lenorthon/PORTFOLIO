CREATE TABLE executions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID REFERENCES workflows(id) ON DELETE CASCADE,
    org_id UUID REFERENCES orgs(id) ON DELETE CASCADE,
    status TEXT DEFAULT 'queued',
    input JSONB,
    output JSONB,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE nodes_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    execution_id UUID REFERENCES executions(id) ON DELETE CASCADE,
    node_id TEXT NOT NULL,
    status TEXT DEFAULT 'queued',
    input JSONB,
    output JSONB,
    error TEXT,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_executions_workflow_id ON executions(workflow_id);
CREATE INDEX idx_nodes_logs_execution_id ON nodes_logs(execution_id); 