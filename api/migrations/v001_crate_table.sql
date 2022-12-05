
CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    email VARCHAR(200),
    name VARCHAR(30),
    password VARCHAR(255),
    PRIMARY KEY(id,email)
);

CREATE TABLE IF NOT EXISTS items (
    id uuid PRIMARY KEY,
    value varchar(255) NOT NULL,
    author varchar(255) NOT NULL
);