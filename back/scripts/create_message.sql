DROP TABLE IF EXISTS messages;

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    userName TEXT NOT NULL,
    message TEXT NOT NULL,
    room TEXT NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL
); 