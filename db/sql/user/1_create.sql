CREATE TABLE app_user (
	id bigserial PRIMARY KEY,
	username VARCHAR(255) UNIQUE,
	full_name VARCHAR(255),
	password VARCHAR (255)
);