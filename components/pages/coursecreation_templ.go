// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"constructor/components"
)

type CourseModule struct {
	Title       string
	Description string
	Lessons     []Lesson
}

type Lesson struct {
	Title       string
	Content     string
	VideoURL    string
	ResourceURL string
	Type        string // video, text, quiz etc
}

func CourseCreation() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen bg-white dark:bg-zinc-900\"><nav class=\"fixed top-0 w-full bg-white/90 dark:bg-zinc-900/90 backdrop-blur-sm border-b border-zinc-200 dark:border-zinc-800 z-50\"><div class=\"max-w-6xl mx-auto px-4 sm:px-6 lg:px-8\"><div class=\"flex justify-between items-center h-16\"><div class=\"flex items-center space-x-4\"><a href=\"/\" class=\"flex items-center space-x-3\"><div class=\"w-9 h-9 bg-black dark:bg-white rounded-lg flex items-center justify-center\"><span class=\"text-white dark:text-black text-lg font-bold\">C</span></div><span class=\"text-lg font-bold text-zinc-900 dark:text-zinc-50\">Constructor</span></a></div></div></div></nav><main class=\"pt-24 pb-16 px-4 sm:px-6 lg:px-8 max-w-6xl mx-auto\"><header class=\"mb-12\"><div class=\"flex justify-between items-center mb-12\"><div><h1 class=\"text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-3\">Create New Course</h1><p class=\"text-base text-zinc-600 dark:text-zinc-400\">Design your comprehensive learning experience</p></div></div></header><form hx-post=\"/api/courses/create\" hx-target=\"#form-result\" class=\"space-y-8 bg-white dark:bg-zinc-800 p-8 rounded-lg border border-zinc-200 dark:border-zinc-700\"><div class=\"grid grid-cols-1 md:grid-cols-2 gap-6\"><div><label for=\"title\" class=\"block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2\">Course Title</label> <input type=\"text\" name=\"title\" required class=\"w-full px-4 py-2 rounded-lg input-style\"></div><div><label for=\"difficulty_level\" class=\"block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2\">Difficulty Level</label> <select name=\"difficulty_level\" required class=\"w-full px-4 py-2 rounded-lg input-style\"><option value=\"\">Select Difficulty</option> <option value=\"beginner\">Beginner</option> <option value=\"intermediate\">Intermediate</option> <option value=\"advanced\">Advanced</option></select></div></div><div><label for=\"description\" class=\"block text-sm font-medium text-zinc-700 dark:text-zinc-300 mb-2\">Course Description</label> <textarea name=\"description\" rows=\"4\" required class=\"w-full px-4 py-2 rounded-lg input-style\"></textarea></div><div id=\"course-modules\"><div class=\"flex justify-between items-center mb-4\"><h2 class=\"text-xl font-semibold\">Course Modules</h2><button type=\"button\" onclick=\"addModule()\" class=\"px-4 py-2 bg-green-500 text-white rounded-lg\">Add Module</button></div><div id=\"modules-container\"></div></div><div class=\"flex justify-end mt-6\"><button type=\"submit\" class=\"px-6 py-2.5 bg-black dark:bg-white text-white dark:text-black rounded-lg\">Create Full Course</button></div></form></main><script>\r\n\t\t\tlet moduleCounter = 0;\r\n\r\n\t\t\tfunction addModule() {\r\n\t\t\t\tmoduleCounter++;\r\n\t\t\t\tconst container = document.getElementById('modules-container');\r\n\t\t\t\tconst moduleHTML = `\r\n\t\t\t\t\t<div class=\"module-block border p-4 mb-4 rounded-lg\" data-module-id=\"${moduleCounter}\">\r\n\t\t\t\t\t\t<div class=\"grid grid-cols-1 md:grid-cols-2 gap-4 mb-4\">\r\n\t\t\t\t\t\t\t<input \r\n\t\t\t\t\t\t\t\ttype=\"text\" \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleCounter}][title]\" \r\n\t\t\t\t\t\t\t\tplaceholder=\"Module Title\" \r\n\t\t\t\t\t\t\t\trequired \r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t/>\r\n\t\t\t\t\t\t\t<textarea \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleCounter}][description]\" \r\n\t\t\t\t\t\t\t\tplaceholder=\"Module Description\" \r\n\t\t\t\t\t\t\t\trows=\"2\"\r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t></textarea>\r\n\t\t\t\t\t\t</div>\r\n\t\t\t\t\t\t\r\n\t\t\t\t\t\t<div id=\"lessons-container-${moduleCounter}\">\r\n\t\t\t\t\t\t</div>\r\n\r\n\t\t\t\t\t\t<div class=\"flex justify-between mt-4\">\r\n\t\t\t\t\t\t\t<button \r\n\t\t\t\t\t\t\t\ttype=\"button\"\r\n\t\t\t\t\t\t\t\tonclick=\"addLesson(${moduleCounter})\"\r\n\t\t\t\t\t\t\t\tclass=\"px-4 py-2 bg-blue-500 text-white rounded-lg\"\r\n\t\t\t\t\t\t\t>\r\n\t\t\t\t\t\t\t\tAdd Lesson\r\n\t\t\t\t\t\t\t</button>\r\n\t\t\t\t\t\t\t<button \r\n\t\t\t\t\t\t\t\ttype=\"button\"\r\n\t\t\t\t\t\t\t\tonclick=\"removeModule(${moduleCounter})\"\r\n\t\t\t\t\t\t\t\tclass=\"px-4 py-2 bg-red-500 text-white rounded-lg\"\r\n\t\t\t\t\t\t\t>\r\n\t\t\t\t\t\t\t\tRemove Module\r\n\t\t\t\t\t\t\t</button>\r\n\t\t\t\t\t\t</div>\r\n\t\t\t\t\t</div>\r\n\t\t\t\t`;\r\n\t\t\t\tcontainer.insertAdjacentHTML('beforeend', moduleHTML);\r\n\t\t\t}\r\n\r\n\t\t\tfunction addLesson(moduleId) {\r\n\t\t\t\tconst container = document.getElementById(`lessons-container-${moduleId}`);\r\n\t\t\t\tconst lessonCounter = container.children.length + 1;\r\n\t\t\t\tconst lessonHTML = `\r\n\t\t\t\t\t<div class=\"lesson-block border p-3 mb-3 rounded-lg\">\r\n\t\t\t\t\t\t<div class=\"grid grid-cols-1 md:grid-cols-2 gap-4 mb-3\">\r\n\t\t\t\t\t\t\t<input \r\n\t\t\t\t\t\t\t\ttype=\"text\" \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleId}][lessons][${lessonCounter}][title]\" \r\n\t\t\t\t\t\t\t\tplaceholder=\"Lesson Title\" \r\n\t\t\t\t\t\t\t\trequired \r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t/>\r\n\t\t\t\t\t\t\t<select \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleId}][lessons][${lessonCounter}][type]\"\r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t>\r\n\t\t\t\t\t\t\t\t<option value=\"text\">Text</option>\r\n\t\t\t\t\t\t\t\t<option value=\"video\">Video</option>\r\n\t\t\t\t\t\t\t\t<option value=\"quiz\">Quiz</option>\r\n\t\t\t\t\t\t\t</select>\r\n\t\t\t\t\t\t</div>\r\n\t\t\t\t\t\t\r\n\t\t\t\t\t\t<textarea \r\n\t\t\t\t\t\t\tname=\"modules[${moduleId}][lessons][${lessonCounter}][content]\" \r\n\t\t\t\t\t\t\tplaceholder=\"Lesson Content\" \r\n\t\t\t\t\t\t\trows=\"3\"\r\n\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style mb-3\"\r\n\t\t\t\t\t\t></textarea>\r\n\r\n\t\t\t\t\t\t<div class=\"grid grid-cols-1 md:grid-cols-2 gap-4\">\r\n\t\t\t\t\t\t\t<input \r\n\t\t\t\t\t\t\t\ttype=\"text\" \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleId}][lessons][${lessonCounter}][video_url]\" \r\n\t\t\t\t\t\t\t\tplaceholder=\"Video URL (optional)\" \r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t/>\r\n\t\t\t\t\t\t\t<input \r\n\t\t\t\t\t\t\t\ttype=\"text\" \r\n\t\t\t\t\t\t\t\tname=\"modules[${moduleId}][lessons][${lessonCounter}][resource_url]\" \r\n\t\t\t\t\t\t\t\tplaceholder=\"Resource URL (optional)\" \r\n\t\t\t\t\t\t\t\tclass=\"w-full px-4 py-2 rounded-lg input-style\"\r\n\t\t\t\t\t\t\t/>\r\n\t\t\t\t\t\t</div>\r\n\t\t\t\t\t</div>\r\n\t\t\t\t`;\r\n\t\t\t\tcontainer.insertAdjacentHTML('beforeend', lessonHTML);\r\n\t\t\t}\r\n\r\n\t\t\tfunction removeModule(moduleId) {\r\n\t\t\t\tconst moduleToRemove = document.querySelector(`[data-module-id=\"${moduleId}\"]`);\r\n\t\t\t\tif (moduleToRemove) {\r\n\t\t\t\t\tmoduleToRemove.remove();\r\n\t\t\t\t}\r\n\t\t\t}\r\n\t\t</script></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CourseCreationBase(cmp templ.Component) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var3 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = cmp.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = components.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var3), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
