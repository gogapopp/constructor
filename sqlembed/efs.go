package sqlembed

import _ "embed"

// Embed SQL files

//go:embed courses.sql
var CoursesTable string

//go:embed course_access.sql
var CourseAccessTable string

//go:embed course_modules.sql
var CourseModulesTable string

//go:embed course_lessons.sql
var CourseLessonsTable string

//go:embed users.sql
var UsersTable string
