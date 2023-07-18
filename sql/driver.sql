CREATE TABLE users (
    user_id numeric(10, 0) NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    age numeric(3, 0) NOT NULL,
    phone character varying(20) COLLATE pg_catalog."default",
    CONSTRAINT user_table_pkey PRIMARY KEY (user_id)
);

INSERT INTO users (user_id, name, age, phone) VALUES (3, 'Jenny', 34, NULL);
INSERT INTO users (user_id, name, age, phone) VALUES (2, 'Tom', 29, '1 800 123 1234');
INSERT INTO users (user_id, name, age, phone) VALUES (1, 'John', 28, NULL);
 
