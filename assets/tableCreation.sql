CREATE TABLE users (
	id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    lastName VARCHAR(45) NOT NULL,
    place_of_birth VARCHAR(200) NOT NULL,
    date_of_birth VARCHAR(200) NOT NULL,
    email VARCHAR(45) NOT NULL
);

INSERT INTO users (name, lastName, place_of_birth, date_of_birth, email) VALUES ("Mark", "Aboian", "Ciudad Autonoma de Buenos Aires, Argentina", "28/01/2002", "abomark09@gmail.com");