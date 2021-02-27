CREATE TABLE public.user_auth
(
    user_id character varying(24),
    user_name character varying(16),
    password character varying(60),
    email_address text,
    member_since timestamp without time zone,
    last_login timestamp without time zone,
    PRIMARY KEY (user_id)
);