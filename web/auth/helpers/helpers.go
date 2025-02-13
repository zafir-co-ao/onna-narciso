package helpers

import "github.com/a-h/templ"

func TriggerEvent(target, event string) templ.ComponentScript {
	return templ.JSFuncCall("htmx.trigger", target, event)
}
