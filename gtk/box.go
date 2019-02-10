package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/whakapapa/gtkgo/glib"
)

//gtk_box_bar_set_center_widget(GtkBox *box,GtkWidget *center_widget)
func (a *Box) SetCenterWidget(child IWidget) {
	if child == nil {
		C.gtk_box_set_center_widget(a.native(), nil)
		} else {
			C.gtk_box_set_center_widget(a.native(), child.toWidget())
		}
	}

	//gtk_box_bar_get_center_widget(GtkBox *box)
	func (a *Box) GetCenterWidget() *Widget {
		w := C.gtk_box_get_center_widget(a.native())
		if w == nil {
			return nil
		}
		return &Widget{glib.InitiallyUnowned{glib.Take(unsafe.Pointer(w))}}
	}
