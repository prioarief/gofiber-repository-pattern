CREATE TABLE books (
	id INT auto_increment NOT NULL,
	title varchar(100) NULL,
	description TEXT NULL,
	price DOUBLE NULL,
	CONSTRAINT books_PK PRIMARY KEY (id)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_0900_ai_ci;
