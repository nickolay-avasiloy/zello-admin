-- Drop the GIN index
DROP INDEX IF EXISTS idx_users_tsv;

-- Drop the ts_vector column
ALTER TABLE users DROP COLUMN IF EXISTS tsv;

-- Drop the trigger
DROP TRIGGER IF EXISTS tsvectorupdate ON users;
