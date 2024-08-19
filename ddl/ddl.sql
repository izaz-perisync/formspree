-- public.projects definition

-- Drop table

-- DROP TABLE public.projects;

CREATE TABLE public.projects (
	id bigserial NOT NULL,
	project_name text NULL,
	project_type text NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	created_by int8 NULL,
	timezone text DEFAULT ''::text NULL,
	restricted_domain text DEFAULT ''::text NULL,
	template_ids _int8 DEFAULT '{}'::bigint[] NULL,
	CONSTRAINT projects_pkey PRIMARY KEY (id)
);

-- public.forms definition

-- Drop table

-- DROP TABLE public.forms;

CREATE TABLE public.forms (
	id bigserial NOT NULL,
	project_id int8 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	user_id int8 NULL,
	"name" text NULL,
	linked_email_id text NULL,
	form_key text NULL,
	"enable" bool DEFAULT true NULL,
	"reCAPTCHA" bool DEFAULT false NULL,
	redirect_url text DEFAULT ''::text NULL,
	store_submissions bool DEFAULT true NULL
);


-- public.forms foreign keys

ALTER TABLE public.forms ADD CONSTRAINT forms_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;