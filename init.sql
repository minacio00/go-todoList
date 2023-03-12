CREATE TABLE IF NOT EXISTS lul (
    id SERIAL PRIMARY KEY,
    name TEXT
);
CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    name TEXT,
    description TEXT
);
INSERT INTO lul (
	name
)
VALUES ('LUCAS'),
('GABRIEL'),
('JOÃO');

INSERT INTO todos (name, description) VALUES ('eat', 'i need to eat today'), ('Study','today, i''m going to study go')