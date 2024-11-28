package course

import (
	"constructor/components/pages"
	"constructor/internal/http-server/middlewares"
	"constructor/internal/lib/render"
	"constructor/internal/model"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

func CourseCreationPage(logger *zap.SugaredLogger, courseService CourseService) http.HandlerFunc {
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

			// extract creator ID from session/auth context (replace with actual auth logic)
			creatorID, ok := getCreatorIDFromContext(ctx)
			if !ok {
				logger.Error("No creator ID found in context")
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			creatorIDint, err := strconv.ParseInt(creatorID, 16, 32)
			if err != nil {
				logger.Errorf("Parse int error (create course handler): %v", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}

			// create course from form data
			course := model.Course{
				Title:           r.Form.Get("title"),
				Description:     r.Form.Get("description"),
				DifficultyLevel: r.Form.Get("difficulty_level"),
				CreatedAt:       time.Now(),
				CreatorID:       int(creatorIDint),
				Modules:         parseModules(r.Form),
			}

			// attempt to save course
			if err := courseService.CreateCourse(ctx, course); err != nil {
				logger.Errorf("Course creation failed: %v", err)
				msg := "Failed to create course"
				renderError(ctx, w, logger, op, msg, pages.CourseCreationBase, pages.CourseCreation)
				return
			}

			// redirect based on request type
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
func getCreatorIDFromContext(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(middlewares.UserIDKey).(string)
	if !ok {
		return "", false
	}
	return user, true
}

// parseModules extracts modules and lessons from form data
func parseModules(form url.Values) []model.Module {
	modules := []model.Module{}

	// iterate through potential modules
	for i := 1; ; i++ {
		moduleKey := fmt.Sprintf("modules[%d]", i)

		// check if module exists
		titleKey := fmt.Sprintf("%s[title]", moduleKey)
		if form.Get(titleKey) == "" {
			break // No more modules
		}

		// create module
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

	// iterate through potential lessons in this module
	for j := 1; ; j++ {
		lessonKey := fmt.Sprintf("modules[%d][lessons][%d]", moduleIndex, j)

		// check if lesson exists
		titleKey := fmt.Sprintf("%s[title]", lessonKey)
		if form.Get(titleKey) == "" {
			break // No more lessons in this module
		}

		// create lesson
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
