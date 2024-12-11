package model

import "time"

type (
	// SignUpUser the user model for SignInPage handler.
	SignUpUser struct {
		UserID       int       `json:""`
		Email        string    `json:"email" validate:"required"`
		Username     string    `json:"username" validate:"required"`
		PasswordHash string    `json:"password" validate:"required"`
		CreatedAt    time.Time `json:""`
		MetaInfo     string    `json:""`
		Role         string    `json:""`
	}
	// SignInUser the user model for SignInPage handler.
	SignInUser struct {
		Email        string `json:"email"`
		Username     string `json:"username"`
		PasswordHash string `json:"password" validate:"required"`
	}
)

type Course struct {
	ID              int       `json:"id"`
	Title           string    `json:"title" validate:"required"`
	Description     string    `json:"description" validate:"required"`
	CreatorID       int       `json:"creator_id" validate:"required"`
	CreatedAt       time.Time `json:"created_at" validate:"required"`
	DifficultyLevel string    `json:"difficulty_level" validate:"required"`
	Modules         []Module  `json:"modules"`
}

type Module struct {
	ModuleID    int      `json:"module_id"`
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description"`
	Lessons     []Lesson `json:"lessons"`
}

type Lesson struct {
	LessonID    int    `json:"lesson_id"`
	Title       string `json:"title" validate:"required"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	VideoURL    string `json:"video_url"`
	ResourceURL string `json:"resource_url"`
}
