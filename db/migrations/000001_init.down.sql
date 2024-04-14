DROP schema IF EXISTS users CASCADE;

DROP schema IF EXISTS organizations CASCADE;

DROP schema IF EXISTS jobs CASCADE;

DROP TYPE IF EXISTS public.employment_type CASCADE;

DROP TYPE IF EXISTS public.pay_type CASCADE;

DROP FUNCTION IF EXISTS public.updated_timestamp() CASCADE;
