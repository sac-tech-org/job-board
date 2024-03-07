/* 
 extensions
 */
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

/*
 functions
 */
CREATE OR REPLACE FUNCTION updated_timestamp() RETURNS TRIGGER AS $$ BEGIN NEW.updated_at := CURRENT_TIMESTAMP;

RETURN NEW;

END;

$$ LANGUAGE PLPGSQL;

/*
 users schema
 */
CREATE SCHEMA IF NOT EXISTS users;

-- user table
CREATE TABLE IF NOT EXISTS users.user (
  id SERIAL PRIMARY KEY,
  user_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  username VARCHAR(255) NOT NULL UNIQUE,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER users_user_updated_trigger BEFORE
UPDATE ON users.user FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- email table
CREATE TABLE IF NOT EXISTS users.email (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users.user(id) NOT NULL,
  address VARCHAR(255) NOT NULL UNIQUE,
  primary_address BOOLEAN NOT NULL DEFAULT FALSE,
  verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER users_email_updated_trigger BEFORE
UPDATE ON users.email FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

/*
 companies scema
 */
CREATE SCHEMA IF NOT EXISTS companies;

-- company table
CREATE TABLE IF NOT EXISTS companies.company (
  id SERIAL PRIMARY KEY,
  company_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_company_updated_trigger BEFORE
UPDATE ON companies.company FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- representative table
CREATE TABLE IF NOT EXISTS companies.representative (
  id SERIAL PRIMARY KEY,
  representative_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_id INT REFERENCES companies.company(id) NOT NULL,
  user_uuid UUID REFERENCES users.user(user_uuid) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_representative_updated_trigger BEFORE
UPDATE ON companies.representative FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

ALTER TABLE companies.company
ADD COLUMN IF NOT EXISTS created_by INT REFERENCES companies.representative(id) NOT NULL,
  ADD COLUMN IF NOT EXISTS owned_by INT REFERENCES companies.representative(id) NOT NULL;

-- email table
CREATE TABLE IF NOT EXISTS companies.email (
  id SERIAL PRIMARY KEY,
  email_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  representative_id INT REFERENCES companies.representative(id) NOT NULL,
  address VARCHAR(255) NOT NULL UNIQUE,
  primary_address BOOLEAN NOT NULL DEFAULT FALSE,
  verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_email_updated_trigger BEFORE
UPDATE ON companies.email FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- profile table
CREATE TABLE IF NOT EXISTS companies.profile (
  id SERIAL PRIMARY KEY,
  profile_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_uuid UUID REFERENCES companies.company(company_uuid) NOT NULL,
  name VARCHAR(255) NOT NULL,
  description TEXT,
  website_url VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_profile_updated_trigger BEFORE
UPDATE ON companies.profile FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- permissions table
CREATE TABLE IF NOT EXISTS companies.permissions (
  id SERIAL PRIMARY KEY,
  permission_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_uuid UUID REFERENCES companies.company(company_uuid) NOT NULL,
  user_uuid UUID REFERENCES users.user(user_uuid) NOT NULL,
  update_profile BOOLEAN NOT NULL DEFAULT FALSE,
  post_job BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/*
 jobs schema
 */
CREATE SCHEMA IF NOT EXISTS jobs;

-- job table
CREATE TABLE IF NOT EXISTS jobs.job (
  id SERIAL PRIMARY KEY,
  job_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_id INT REFERENCES companies.company(id) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  publish_start_date TIMESTAMP NOT NULL,
  publish_end_date TIMESTAMP NOT NULL,
  publish_override BOOLEAN NOT NULL DEFAULT FALSE,
  deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER jobs_job_updated_trigger BEFORE
UPDATE ON jobs.job FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- permissions table
CREATE TABLE IF NOT EXISTS jobs.job_permissions (
  id SERIAL PRIMARY KEY,
  job_id INT REFERENCES jobs.job(id) NOT NULL UNIQUE,
  user_id INT REFERENCES users.user(id) NOT NULL,
  update_job BOOLEAN NOT NULL DEFAULT FALSE,
  delete_job BOOLEAN NOT NULL DEFAULT FALSE,
  publish BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER jobs_job_permissions_updated_trigger BEFORE
UPDATE ON jobs.job_permissions FOR EACH ROW EXECUTE FUNCTION updated_timestamp();
