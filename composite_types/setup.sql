CREATE TYPE Address AS (
    street TEXT,
    city TEXT
);

CREATE TABLE Person (
    id INT GENERATED ALWAYS AS IDENTITY,
    name TEXT,
    address Address
);

INSERT INTO Person (name, address) VALUES ('Teja', ROW('Skogsangsvagen', 'Stockholm'));

SELECT * FROM Person;