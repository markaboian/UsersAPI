CREATE TABLE users (
	id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(45) NOT NULL,
    lastName VARCHAR(45) NOT NULL,
    place_of_birth VARCHAR(200) NOT NULL,
    date_of_birth VARCHAR(200) NOT NULL,
    email VARCHAR(45) NOT NULL
);

CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (name, lastName, place_of_birth, date_of_birth, email) VALUES ("John", "Doe", "Los Angeles, California", "01/01/1999", "johndoe@gmail.com");
INSERT INTO products (name, price, user_id) VALUES ("Apple IPhone 15 PRO", 1232.00, 1);
