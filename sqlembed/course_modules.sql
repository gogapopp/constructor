CREATE TABLE IF NOT EXISTS course_modules (
    module_id SERIAL PRIMARY KEY,
    course_id INTEGER NOT NULL,
    parent_module_id INTEGER,
    title TEXT NOT NULL,
    description TEXT,
    order_index INTEGER,
    duration_minutes INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES courses(course_id),
    FOREIGN KEY (parent_module_id) REFERENCES course_modules(module_id)
);
