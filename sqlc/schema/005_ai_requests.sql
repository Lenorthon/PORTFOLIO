CREATE TABLE ai_requests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    execution_id UUID REFERENCES executions(id),
    node_id TEXT,
    provider TEXT NOT NULL,
    model TEXT NOT NULL,
    input JSONB,
    output JSONB,
    cost NUMERIC DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_ai_requests_execution_id ON ai_requests(execution_id); 