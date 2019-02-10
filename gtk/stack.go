package gtk

// #include <stdlib.h>
// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/whakapapa/gtkgo/glib"
)

// GetChildByName is a wrapper around gtk_stack_get_child_by_name().
func (v *Stack) GetChildByName(name string) *Widget {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	c := C.gtk_stack_get_child_by_name(v.native(), (*C.gchar)(cstr))
	if c == nil {
		return nil
	}
	return wrapWidget(glib.Take(unsafe.Pointer(c)))
}

// GetTransitionRunning is a wrapper around gtk_stack_get_transition_running().
func (v *Stack) GetTransitionRunning() bool {
	c := C.gtk_stack_get_transition_running(v.native())
	return gobool(c)
}
