# PORTOFOLIO - Full Platform (MVP)

## Structure
- backend/ : Go backend (complete)
- frontend/ : React + Vite (minimal)
- infra/ : Dockerfiles and docker-compose
- backend/internal/db/migrations/ : SQL migrations
- docs/ : documentation

## Quick start (local)
1. Copy `.env.example` to `.env` and adjust.
2. Start DB (docker-compose) or run Postgres locally.
3. Run migrations:
   ```bash
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/001_init.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/002_workflows.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/003_marketing.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/004_analytics.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/005_cms.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/006_integrations.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/007_plans_billing.sql
   psql -h localhost -U postgres -d portofolio -f backend/internal/db/migrations/008_usage_counters.sql
