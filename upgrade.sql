-- 2022-09-13
-- closure_tree

CREATE TABLE public.t_meta (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,  
	master_id uuid NULL,	  
	meta_name varchar(256) NULL,	
	extra jsonb NULL,
	create_user_id uuid NULL,
	create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	update_user_id uuid NULL,
	update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	is_del bool NULL DEFAULT false,
	CONSTRAINT t_meta_pk PRIMARY KEY (id)
);


CREATE TABLE public.t_meta_path (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	tenant_id uuid NULL,	
	ancestor_id uuid NULL,
	descendant_id uuid NULL,
	distance int4,	
	create_user_id uuid NULL,
	create_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	update_user_id uuid NULL,
	update_time timestamp NULL DEFAULT CURRENT_TIMESTAMP,
	is_del bool NULL DEFAULT false,
	CONSTRAINT t_meta_path_pk PRIMARY KEY (id)
);

CREATE UNIQUE INDEX t_ancester_id_descendant_id_distance_idx ON public.t_meta_path (ancestor_id, descendant_id, distance);

