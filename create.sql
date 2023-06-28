CREATE SCHEMA SecuriGroup;
GO

CREATE TABLE SecuriGroup.employees (
	id INT IDENTITY(1,1) PRIMARY KEY NOT NULL,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	username VARCHAR(255),
	password VARCHAR(255),
	email VARCHAR(255),
	date_of_birth DATE,
	department_id INT,
	position VARCHAR(20) NOT NULL CHECK (position IN('Junior', 'Senior', 'Leader')),
);

ALTER TABLE SecuriGroup.employees
ADD CONSTRAINT UC_Username UNIQUE (username);
