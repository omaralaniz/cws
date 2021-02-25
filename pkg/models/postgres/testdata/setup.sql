CREATE TABLE authors(
    id serial,
    name text NOT NULL,
    email text NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created date not null,
    active boolean not null DEFAULT false,
    CONSTRAINT authors_pkey PRIMARY KEY (id),
    CONSTRAINT author_uc_email UNIQUE (email)
);

ALTER TABLE authors OWNER to postgres;


CREATE TABLE posts (
    id serial,
    title varchar(60) NOT NULL,
    author text NOT NULL,
    content text NOT NULL,
    modified date NOT NULL,
    category text NOT NULL,
    summary text NOT NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
);

ALTER TABLE posts OWNER to postgres;

INSERT INTO authors (name, email, hashed_password, created, active) VALUES (

    'Mockaton Smith',
    'mockaton@gmail.com',
    '$2a$12$NuTjWXm3KKntReFwyBVHyuf/to.HEwTy.eS206TNfkGfr6HzGJSWG',
    '2020-12-18',
    'false'
);