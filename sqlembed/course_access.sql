CREATE TABLE IF NOT EXISTS course_access (
    access_id INTEGER PRIMARY KEY AUTOINCREMENT,
    course_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    granted_at TEXT DEFAULT (datetime('now')),
    granted_by INTEGER,
    access_type TEXT DEFAULT 'view',
    valid_from TEXT DEFAULT (datetime('now')),
    valid_until TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    access_status TEXT DEFAULT 'active',
    FOREIGN KEY (course_id) REFERENCES courses(course_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (granted_by) REFERENCES users(user_id),
    UNIQUE(course_id, user_id)
);

-- postgres
-- CREATE TABLE IF NOT EXISTS course_access (
--     access_id SERIAL PRIMARY KEY,
--     course_id INTEGER NOT NULL,
--     user_id INTEGER NOT NULL,
--     granted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     granted_by INTEGER,
--     access_type TEXT DEFAULT 'view',
--     valid_from TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
--     valid_until TIMESTAMP,
--     is_active BOOLEAN DEFAULT TRUE,
--     access_status TEXT DEFAULT 'active',
--     FOREIGN KEY (course_id) REFERENCES courses(course_id),
--     FOREIGN KEY (user_id) REFERENCES users(user_id),
--     FOREIGN KEY (granted_by) REFERENCES users(user_id),
--     UNIQUE(course_id, user_id)
-- );
