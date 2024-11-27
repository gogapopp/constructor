package course

import (
	"constructor/components/pages"
	"constructor/internal/lib/render"
	"constructor/internal/model"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type CourseService interface {
	SignUp(ctx context.Context, user model.SignUpUser) error
	SignIn(ctx context.Context, user model.SignInUser) (string, error)
}

func CourseCreationPage(logger *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "course.course_create.CourseCreationPage"
		ctx := r.Context()

		switch r.Method {
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				logger.Errorf("Form parse error: %v", err)
				msg := "Failed to parse form data"
				renderError(ctx, w, logger, op, msg, pages.CourseCreationBase, pages.CourseCreation)
				return
			}

			// Create course from form data
			course := model.Course{
				Title:           r.Form.Get("title"),
				Description:     r.Form.Get("description"),
				DifficultyLevel: r.Form.Get("difficulty_level"),
				CreatedAt:       time.Now(),
				// TODO: Set CreatorID from session or authentication
				CreatorID: 1, // Placeholder
			}

			// Process modules dynamically
			course.Modules = parseModules(r.Form)

			// Log the parsed course for debugging
			logger.Infof("Parsed Course: %+v", course)

			// TODO: Save course to database
			// courseID, err := courseService.Create(ctx, course)
			// if err != nil {
			//     // Handle error
			//     return
			// }

			if r.Header.Get("HX-Request") == "true" {
				w.Header().Set("HX-Redirect", "/course/create?success=true")
				w.WriteHeader(http.StatusSeeOther)
				return
			}

			http.Redirect(w, r, "/course/create?success=true", http.StatusSeeOther)
			return

		case http.MethodGet:
			msg := ""
			if r.URL.Query().Get("success") == "true" {
				msg = "Course created successfully"
			}

			if err := render.Render(r.Context(), w, pages.CourseCreationBase(pages.CourseCreation(msg))); err != nil {
				logger.Errorf("Render error: %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	}
}

// parseModules extracts modules and lessons from form data
func parseModules(form url.Values) []model.Module {
	modules := []model.Module{}

	// Iterate through potential modules
	for i := 1; ; i++ {
		moduleKey := fmt.Sprintf("modules[%d]", i)

		// Check if module exists
		titleKey := fmt.Sprintf("%s[title]", moduleKey)
		if form.Get(titleKey) == "" {
			break // No more modules
		}

		// Create module
		module := model.Module{
			Title:       form.Get(titleKey),
			Description: form.Get(fmt.Sprintf("%s[description]", moduleKey)),
			Lessons:     parseLessons(form, i),
		}

		modules = append(modules, module)
	}

	return modules
}

// parseLessons extracts lessons for a specific module
func parseLessons(form url.Values, moduleIndex int) []model.Lesson {
	lessons := []model.Lesson{}

	// Iterate through potential lessons in this module
	for j := 1; ; j++ {
		lessonKey := fmt.Sprintf("modules[%d][lessons][%d]", moduleIndex, j)

		// Check if lesson exists
		titleKey := fmt.Sprintf("%s[title]", lessonKey)
		if form.Get(titleKey) == "" {
			break // No more lessons in this module
		}

		// Create lesson
		lesson := model.Lesson{
			Title:       form.Get(titleKey),
			Type:        form.Get(fmt.Sprintf("%s[type]", lessonKey)),
			Content:     form.Get(fmt.Sprintf("%s[content]", lessonKey)),
			VideoURL:    form.Get(fmt.Sprintf("%s[video_url]", lessonKey)),
			ResourceURL: form.Get(fmt.Sprintf("%s[resource_url]", lessonKey)),
		}

		lessons = append(lessons, lesson)
	}

	return lessons
}

func renderError(ctx context.Context, w http.ResponseWriter, logger *zap.SugaredLogger, op, msg string, basePage func(templ.Component) templ.Component, authPage func(string) templ.Component) {
	w.Header().Set("Content-Type", "text/html")
	if err := render.Render(ctx, w, basePage(authPage(msg))); err != nil {
		logger.Errorf("[%s] %s: %w", middleware.GetReqID(ctx), op, err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
