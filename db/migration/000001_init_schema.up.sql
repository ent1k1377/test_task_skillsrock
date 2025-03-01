CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT,
    description TEXT,
    status TEXT,
    created_at TIMESTAMP,
    update_at TIMESTAMP
);