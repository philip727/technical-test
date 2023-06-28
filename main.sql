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

USE SecuriGroup;
INSERT INTO SecuriGroup.employees (first_name, last_name, password, email, date_of_birth, department_id, position)
VALUES ('John', 'Doe', 'randompassword', 'johndoe@gmail.com', '2002-11-25', 12, 'Junior');

USE SecuriGroup;
INSERT INTO SecuriGroup.employees (first_name, last_name, password, email, date_of_birth, department_id, position)
VALUES ('John', 'Davids', 'randompassword', 'johndoe@gmail.com', '1994-11-25', 12, 'Senior');


USE SecuriGroup;
SELECT * FROM SecuriGroup.employees;

UPDATE SecuriGroup.employees
SET department_id = '2'
WHERE first_name = 'John' and last_name = 'Davids';