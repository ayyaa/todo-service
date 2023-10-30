/**
  This is the SQL script that will be used to initialize the database schema.
  We will evaluate you based on how well you design your database.
  1. How you design the tables.
  2. How you choose the data types and keys.
  3. How you name the fields.
  In this assignment we will use PostgreSQL as the database.
  */

/** This is test table. Remove this table and replace with your own tables. */
-- Membuat tabel "list"
CREATE TABLE list (
    id serial PRIMARY KEY,
    title varchar(100) NOT NULL,
    description text NOT NULL,
    priority int NOT NULL,
    parent_id int,
    status varchar(10) NULL DEFAULT 'active'::character varying,
    updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES list (id)
);

-- public.attachment definition

-- Drop table

-- DROP TABLE public.attachment;

CREATE TABLE public.attachment (
	id serial NOT NULL PRIMARY key,
	list_id int4 NULL,
	filepath varchar(255) NOT NULL,
	filename varchar(255) NOT NULL,
	status varchar(10) NULL DEFAULT 'active'::character varying,
  updated_at timestamp DEFAULT CURRENT_TIMESTAMP,
  created_at timestamp DEFAULT CURRENT_TIMESTAMP,
	FOREIGN KEY (list_id) REFERENCES list (id)
);
