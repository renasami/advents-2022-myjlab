
CREATE TABLE IF NOT EXISTS users {
    id uuid PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(30)

}

CREATE TABLE IF NOT EXISTS items {
    id uuid PRIMARY KEY,
    value varchar(255) NOT NULL,
    author varchar(255) NOT NULL
}