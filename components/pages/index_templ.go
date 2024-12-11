// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "constructor/components"
import "constructor/components/partials"

func Home(title string) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"min-h-screen bg-white dark:bg-zinc-900\"><!-- Navigation --><nav class=\"fixed top-0 w-full bg-white/90 dark:bg-zinc-900/90 backdrop-blur-sm border-b border-zinc-200 dark:border-zinc-800 z-50\"><div class=\"max-w-6xl mx-auto px-4 sm:px-6 lg:px-8\"><div class=\"flex justify-between items-center h-16\"><div class=\"flex items-center space-x-4\"><a href=\"/\" class=\"flex items-center space-x-3\"><div class=\"w-9 h-9 bg-black dark:bg-white rounded-lg flex items-center justify-center\"><span class=\"text-white dark:text-black text-lg font-bold\">C</span></div><span class=\"text-lg font-bold text-zinc-900 dark:text-zinc-50\">Constructor</span></a></div><div class=\"flex items-center space-x-6\"><a href=\"/signin\" class=\"px-4 py-2 bg-black dark:bg-white text-white dark:text-black rounded-lg font-medium hover:bg-zinc-800 dark:hover:bg-zinc-100 transform transition-all duration-200 hover:scale-[1.02] active:scale-[0.98]\">Логин</a> <a href=\"/signup\" class=\"px-4 py-2 bg-black dark:bg-white text-white dark:text-black rounded-lg font-medium hover:bg-zinc-800 dark:hover:bg-zinc-100 transform transition-all duration-200 hover:scale-[1.02] active:scale-[0.98]\">Регистрация</a></div></div></div></nav><!-- Main Content --><main class=\"pt-24 pb-16 px-4 sm:px-6 lg:px-8 max-w-6xl mx-auto\"><header class=\"mb-12\"><div class=\"flex justify-between items-center mb-12\"><div><h1 class=\"text-4xl font-bold text-zinc-900 dark:text-zinc-50 mb-3\">Добро пожаловать в Constructor</h1><p class=\"text-base text-zinc-600 dark:text-zinc-400\">Ваша идеальная платформа для создания курсов и управления ими</p></div></div></header><!-- Project Information --><section class=\"mb-12\"><h2 class=\"text-2xl font-bold text-zinc-900 dark:text-zinc-50 mb-6\">О проекте</h2><p class=\"text-base text-zinc-600 dark:text-zinc-400 mb-4\">Constructor - это мощный инструмент, который поможет вам с легкостью создавать свои курсы, управлять ими и развертывать их. Независимо от того, являетесь ли вы разработчиком, дизайнером или руководителем проекта.</p><p class=\"text-base text-zinc-600 dark:text-zinc-400 mb-4\">Constructor, может:</p><ul class=\"list-disc list-inside text-zinc-600 dark:text-zinc-400 mb-4\"><li>Создавайте курсы и управляйте ими без особых усилий</li><li>Развертывайте свои курсы одним щелчком мыши</li></ul></section><!-- Call to Action --><section class=\"text-center\"><h2 class=\"text-2xl font-bold text-zinc-900 dark:text-zinc-50 mb-6\">Начните уже сегодня</h2><a href=\"/constructor\" class=\"px-6 py-2.5 bg-black dark:bg-white text-white dark:text-black rounded-lg font-medium hover:bg-zinc-800 dark:hover:bg-zinc-100 transform transition-all duration-200 hover:scale-[1.02] active:scale-[0.98]\">Начать</a></section></main></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = partials.Footer().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func HomeBase(cmp templ.Component) templ.Component {
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
