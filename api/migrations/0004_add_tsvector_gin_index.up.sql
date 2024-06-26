-- Add a new column for ts_vector
ALTER TABLE users ADD COLUMN tsv tsvector;

-- Update the ts_vector column with concatenated text
UPDATE users SET tsv =
  setweight(to_tsvector(coalesce(username, '')), 'A') ||
  setweight(to_tsvector(coalesce(company, '')), 'B') ||
  setweight(to_tsvector(coalesce(phone_number, '')), 'C');

-- Create a trigger to update the ts_vector column on data change
CREATE TRIGGER tsvectorupdate BEFORE INSERT OR UPDATE
  ON users FOR EACH ROW EXECUTE FUNCTION
  tsvector_update_trigger(tsv, 'pg_catalog.english', username, company, phone_number);

-- Create GIN index on the ts_vector column
CREATE INDEX idx_users_tsv ON users USING gin(tsv);
