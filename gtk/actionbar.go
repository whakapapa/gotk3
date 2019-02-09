package gtk

// #include <gtk/gtk.h>
// #include "actionbar.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_action_bar_get_type()), marshalActionBar},
	}

	glib.RegisterGValueMarshalers(tm)

	WrapMap["GtkActionBar"] = wrapActionBar
}

//GtkActionBar
type ActionBar struct {
	Bin
}

func (v *ActionBar) native() *C.GtkActionBar {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkActionBar(p)
}

func marshalActionBar(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapActionBar(glib.Take(unsafe.Pointer(c))), nil
}

func wrapActionBar(obj *glib.Object) *ActionBar {
	return &ActionBar{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

//gtk_action_bar_new()
func ActionBarNew() (*ActionBar, error) {
	c := C.gtk_action_bar_new()
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapActionBar(glib.Take(unsafe.Pointer(c))), nil
}

//gtk_action_bar_pack_start(GtkActionBar *action_bar,GtkWidget *child)
func (a *ActionBar) PackStart(child IWidget) {
	C.gtk_action_bar_pack_start(a.native(), child.toWidget())
}

//gtk_action_bar_pack_end(GtkActionBar *action_bar,GtkWidget *child)
func (a *ActionBar) PackEnd(child IWidget) {
	C.gtk_action_bar_pack_end(a.native(), child.toWidget())
}

//gtk_action_bar_set_center_widget(GtkActionBar *action_bar,GtkWidget *center_widget)
func (a *ActionBar) SetCenterWidget(child IWidget) {
	if child == nil {
		C.gtk_action_bar_set_center_widget(a.native(), nil)
		} else {
			C.gtk_action_bar_set_center_widget(a.native(), child.toWidget())
		}
	}

	//gtk_action_bar_get_center_widget(GtkActionBar *action_bar)
	func (a *ActionBar) GetCenterWidget() *Widget {
		w := C.gtk_action_bar_get_center_widget(a.native())
		if w == nil {
			return nil
		}
		return &Widget{glib.InitiallyUnowned{glib.Take(unsafe.Pointer(w))}}
	}
