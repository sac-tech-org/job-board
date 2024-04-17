-- change organizations.permissions relation from user to representative
ALTER TABLE organizations.permissions DROP CONSTRAINT IF EXISTS permissions_user_id_fkey,
  DROP COLUMN IF EXISTS user_id,
  ADD COLUMN IF NOT EXISTS representative_id INT REFERENCES organizations.representative(representative_id) NOT NULL;

-- change jobs.job_permissions relation from user to representative
ALTER TABLE jobs.job_permissions DROP CONSTRAINT IF EXISTS job_permissions_user_id_fkey,
  DROP COLUMN IF EXISTS user_id,
  ADD COLUMN IF NOT EXISTS representative_id INT REFERENCES organizations.representative(representative_id) NOT NULL;

-- rename jobs.job_permissions to jobs.permissions
ALTER TABLE jobs.job_permissions
  RENAME TO permissions;
