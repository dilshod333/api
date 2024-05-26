-- CREATE TABLE users (
--     id SERIAL PRIMARY KEY,
--     name VARCHAR(100) NOT NULL,
--     email VARCHAR(100) NOT NULL
-- );

-- CREATE TABLE orders (
--     id SERIAL PRIMARY KEY,
--     user_id INT NOT NULL,
--     product VARCHAR(100) NOT NULL,
--     amount INT NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(id)
-- );

CREATE TABLE Authors (
     author_id INT PRIMARY KEY,
     name VARCHAR(100) NOT NULL,
     birth_date DATE NOT NULL,
     biography TEXT,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE Books (
     book_id INT PRIMARY KEY,
     title VARCHAR(100) NOT NULL,
     author_id INT REFERENCES Authors(author_id) ON DELETE CASCADE,
     publication_date DATE,
     isbn VARCHAR(20),
     description TEXT,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);