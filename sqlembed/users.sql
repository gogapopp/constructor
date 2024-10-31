CREATE TABLE IF NOT EXISTS users (
	user_id SERIAL PRIMARY KEY,
	email TEXT,
	username TEXT,
	password_hash TEXT,
	created_at TIMESTAMP,
	role TEXT,
	metainfo TEXT
);
CREATE UNIQUE INDEX IF NOT EXISTS username_idx ON users(username);