CREATE TABLE IF NOT EXISTS course_lessons (
	lesson_id INTEGER PRIMARY KEY AUTOINCREMENT,
	module_id INTEGER NOT NULL,
	title TEXT NOT NULL,
	type TEXT,
	content TEXT,
	video_url TEXT,
	resource_url TEXT,
	FOREIGN KEY (module_id) REFERENCES course_modules(module_id)
);