CREATE TABLE users (
	id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    lastName VARCHAR(45) NOT NULL,
    place_of_birth VARCHAR(200) NOT NULL,
    date_of_birth VARCHAR(200) NOT NULL,
    email VARCHAR(45) NOT NULL
);

INSERT INTO users (name, lastName, place_of_birth, date_of_birth, email) VALUES ("John", "Doe", "Los Angeles, California", "01/01/1999", "johndoe@gmail.com");
