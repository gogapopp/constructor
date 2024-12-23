CREATE TABLE IF NOT EXISTS courses (
    course_id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT,
    creator_id INTEGER NOT NULL,
    created_at TEXT DEFAULT (datetime('now')),
    updated_at TEXT,
    status TEXT DEFAULT 'draft',
    difficulty_level TEXT,
    price REAL DEFAULT 0,
    parent_course_id INTEGER,
    metadata TEXT,
    FOREIGN KEY (parent_course_id) REFERENCES courses(course_id)
);

-- postgres
-- CREATE TABLE IF NOT EXISTS courses (
--     course_id SERIAL PRIMARY KEY,
--     title TEXT NOT NULL,
--     description TEXT,
--     creator_id INTEGER NOT NULL,
--     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP,
--     status TEXT DEFAULT 'draft',
--     difficulty_level TEXT,
--     price NUMERIC(10, 2) DEFAULT 0,
--     cover_image_url TEXT,
--     parent_course_id INTEGER,
--     metadata JSONB,
--     FOREIGN KEY (parent_course_id) REFERENCES courses(course_id)
-- );