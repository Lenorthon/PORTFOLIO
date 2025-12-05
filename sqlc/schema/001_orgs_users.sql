-- orgs
CREATE TABLE orgs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    plan TEXT DEFAULT 'free',
    created_at TIMESTAMPTZ DEFAULT now()
);

-- users
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    org_id UUID REFERENCES orgs(id) ON DELETE CASCADE,
    email TEXT UNIQUE NOT NULL,
    name TEXT,
    password_hash TEXT NOT NULL,
    role TEXT DEFAULT 'member',
    created_at TIMESTAMPTZ DEFAULT now()
);
CREATE INDEX idx_users_org_id ON users(org_id);
