CREATE TABLE ArrayExample (
    id INT GENERATED ALWAYS AS IDENTITY,
    name TEXT,
    scores INT[]
);

INSERT INTO ArrayExample (name, scores) VALUES ('Teja', ARRAY[100, 95, 90]);

SELECT * FROM ArrayExample;