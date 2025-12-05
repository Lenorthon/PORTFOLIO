CREATE TABLE credits_ledger (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID REFERENCES orgs(id),
    execution_id UUID REFERENCES executions(id),
    amount NUMERIC NOT NULL,
    reason TEXT,
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE TABLE orgs_credits (
    org_id UUID PRIMARY KEY REFERENCES orgs(id),
    total_credits NUMERIC DEFAULT 0,
    used_credits NUMERIC DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now()
); 