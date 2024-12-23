// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"constructor/components"
	"constructor/internal/model"
	"fmt"
)

func Constructor(courses []model.Course, urls []string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen bg-white dark:bg-zinc-900\"><!-- Navigation --><nav class=\"fixed top-0 w-full bg-white/90 dark:bg-zinc-900/90 backdrop-blur-sm border-b border-zinc-200 dark:border-zinc-800 z-50\"><div class=\"max-w-6xl mx-auto px-4 sm:px-6 lg:px-8\"><div class=\"flex justify-between items-center h-16\"><div class=\"flex items-center space-x-4\"><a href=\"/\" class=\"flex items-center space-x-3\"><div class=\"w-9 h-9 bg-black dark:bg-white rounded-lg flex items-center justify-center\"><span class=\"text-white dark:text-black text-lg font-bold\">C</span></div><span class=\"text-lg font-bold text-zinc-900 dark:text-zinc-50\">Constructor</span></a></div></div></div></nav><!-- Main Content --><main class=\"pt-24 pb-16 px-4 sm:px-6 lg:px-8 max-w-6xl mx-auto\"><header class=\"mb-12\"><div class=\"flex justify-between items-center mb-12\"><div><h1 class=\"text-3xl font-bold text-zinc-900 dark:text-zinc-50 mb-3\">Доступные курсы</h1><p class=\"text-base text-zinc-600 dark:text-zinc-400\">Начните учиться уже сегодня</p></div><a href=\"/course/create\" class=\"px-6 py-2.5 bg-black dark:bg-white text-white dark:text-black\r\n\t\t\t\t\t\t\t  rounded-lg font-medium\r\n\t\t\t\t\t\t\t  hover:bg-zinc-800 dark:hover:bg-zinc-100\r\n\t\t\t\t\t\t\t  transform transition-all duration-200\r\n\t\t\t\t\t\t\t  hover:scale-[1.02] active:scale-[0.98]\r\n\t\t\t\t\t\t\t  flex items-center space-x-2\"><svg class=\"w-5 h-5\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> <span>Создать курс</span></a></div></header><!-- Course Grid --><div id=\"courses-grid\" class=\"grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6\" hx-get=\"/courses\" hx-trigger=\"load\" hx-indicator=\"#loading\"><!-- Course Cards -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, course := range courses {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"group bg-white dark:bg-zinc-800 rounded-lg border border-zinc-200 dark:border-zinc-700\r\n\t\t\t\t\t\t\thover:border-black dark:hover:border-white\r\n\t\t\t\t\t\t\ttransition-all duration-200\"><div class=\"p-5\"><div class=\"flex items-center space-x-2 mb-3\"><span class=\"px-2.5 py-1 text-xs font-medium text-zinc-700 dark:text-zinc-300\r\n\t\t\t\t\t\t\t\t   bg-zinc-100 dark:bg-zinc-700 rounded-full\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", course.ID))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/pages/constructor.templ`, Line: 66, Col: 66}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><h3 class=\"text-lg font-semibold text-zinc-900 dark:text-zinc-50 mb-2\r\n\t\t\t\t\t\t\t\t   group-hover:text-black dark:group-hover:text-white\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(course.Title)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/pages/constructor.templ`, Line: 72, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h3><p class=\"text-sm text-zinc-600 dark:text-zinc-400 mb-4\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(course.Description)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/pages/constructor.templ`, Line: 76, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p><div class=\"flex justify-between items-center\"><span class=\"flex items-center space-x-2\"><svg class=\"w-4 h-4 text-zinc-500 dark:text-zinc-400\" fill=\"none\" stroke=\"currentColor\" viewBox=\"0 0 24 24\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M12 6v6m0 0v6m0-6h6m-6 0H6\"></path></svg> <span class=\"text-sm text-zinc-500 dark:text-zinc-400\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", len(course.Modules)))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/pages/constructor.templ`, Line: 85, Col: 102}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" модули </span></span><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL(fmt.Sprintf("course/%d", course.ID))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"px-4 py-2 bg-black dark:bg-white\r\n\t\t\t\t\t\t\t\t\t\ttext-white dark:text-black text-sm font-medium rounded-lg\r\n\t\t\t\t\t\t\t\t\t\thover:bg-zinc-800 dark:hover:bg-zinc-100\r\n\t\t\t\t\t\t\t\t\t\ttransform transition-all duration-200\r\n\t\t\t\t\t\t\t\t\t\thover:scale-[1.02] active:scale-[0.98]\">Начать</a></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><!-- Loading State --><div id=\"loading\" class=\"hidden\"><div class=\"flex justify-center items-center py-12\"><div class=\"animate-spin rounded-full h-8 w-8 border-2 border-black dark:border-white border-t-transparent\"></div></div></div><!-- Load More --></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ConstructorBase(cmp templ.Component) templ.Component {
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
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var8 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
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
		templ_7745c5c3_Err = components.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var8), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
