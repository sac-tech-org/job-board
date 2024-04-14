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
  user_id SERIAL PRIMARY KEY,
  first_name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  username VARCHAR(255) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER users_user_updated_trigger BEFORE
UPDATE ON users.user FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- email table
CREATE TABLE IF NOT EXISTS users.email (
  email_id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users.user(user_id) NOT NULL,
  address VARCHAR(255) NOT NULL UNIQUE,
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
  company_id SERIAL PRIMARY KEY,
  company_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_company_updated_trigger BEFORE
UPDATE ON companies.company FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- representative table
CREATE TABLE IF NOT EXISTS companies.representative (
  representative_id SERIAL PRIMARY KEY,
  representative_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_id INT REFERENCES companies.company(company_id) NOT NULL,
  user_id UUID REFERENCES users.user(user_id) NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_representative_updated_trigger BEFORE
UPDATE ON companies.representative FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

ALTER TABLE companies.company
ADD COLUMN IF NOT EXISTS created_by INT REFERENCES companies.representative(representative_id) NOT NULL,
  ADD COLUMN IF NOT EXISTS owned_by INT REFERENCES companies.representative(representative_id) NOT NULL;

-- email table
CREATE TABLE IF NOT EXISTS companies.email (
  email_id SERIAL PRIMARY KEY,
  email_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  representative_id INT REFERENCES companies.representative(representative_id) NOT NULL,
  address VARCHAR(255) NOT NULL UNIQUE,
  verified BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER companies_email_updated_trigger BEFORE
UPDATE ON companies.email FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

-- profile table
CREATE TABLE IF NOT EXISTS companies.profile (
  profile_id SERIAL PRIMARY KEY,
  company_id UUID REFERENCES companies.company(company_id) NOT NULL,
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
  permissions_id SERIAL PRIMARY KEY,
  company_id UUID REFERENCES companies.company(company_id) NOT NULL,
  user_id UUID REFERENCES users.user(user_id) NOT NULL,
  update_profile BOOLEAN NOT NULL DEFAULT FALSE,
  post_job BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/*
 jobs schema
 */
CREATE SCHEMA IF NOT EXISTS jobs;

-- enums
CREATE TYPE employment_type AS ENUM (
  'full-time',
  'part-time',
  'contract',
  'temporary',
  'internship',
  'volunteer',
  'per-diem',
  'other'
);

CREATE TYPE pay_type AS ENUM (
  'hourly',
  'salary',
  'commission',
  'other'
);

-- job table
CREATE TABLE IF NOT EXISTS jobs.job (
  job_id SERIAL PRIMARY KEY,
  job_uuid UUID DEFAULT uuid_generate_v4() NOT NULL UNIQUE,
  company_id INT REFERENCES companies.company(company_id) NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER jobs_job_updated_trigger BEFORE
UPDATE ON jobs.job FOR EACH ROW EXECUTE FUNCTION updated_timestamp();

CREATE TABLE IF NOT EXISTS jobs.job_posting (
  job_posting_id SERIAL PRIMARY KEY,
  job_id INT REFERENCES jobs.job(job_id) NOT NULL,
  employment_type employment_type [] NOT NULL CHECK (employment_type <> '{}'::employment_type []),
  pay_rate NUMERIC(10, 2) CHECK (pay_rate >= 0) NOT NULL DEFAULT 0,
  pay_type pay_type NOT NULL,
  publish_start_date TIMESTAMP NOT NULL,
  publish_end_date TIMESTAMP NOT NULL,
  publish_override BOOLEAN NOT NULL DEFAULT FALSE,
  deleted BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- permissions table
CREATE TABLE IF NOT EXISTS jobs.job_permissions (
  id SERIAL PRIMARY KEY,
  job_id INT REFERENCES jobs.job(job_posting_id) NOT NULL UNIQUE,
  user_id INT REFERENCES users.user(user_id) NOT NULL,
  update_job BOOLEAN NOT NULL DEFAULT FALSE,
  delete_job BOOLEAN NOT NULL DEFAULT FALSE,
  post_job BOOLEAN NOT NULL DEFAULT FALSE,
  publish BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER jobs_job_permissions_updated_trigger BEFORE
UPDATE ON jobs.job_permissions FOR EACH ROW EXECUTE FUNCTION updated_timestamp();
