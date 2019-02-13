package gtk

// #include <gtk/gtk.h>
// #include "popover.go.h"
import "C"
import (
	"unsafe"

	"github.com/whakapapa/gtkgo/glib"
)

func init() {
	tm := []glib.TypeMarshaler{
		{glib.Type(C.gtk_popover_get_type()), marshalPopover},
	}

	glib.RegisterGValueMarshalers(tm)
	WrapMap["GtkPopover"] = wrapPopover
}

//TODO(sjon): Implement GtkPopover
//GtkPopover
type Popover struct {
	Bin
}

func (v *Popover) native() *C.GtkPopover {
	if v == nil || v.GObject == nil {
		return nil
	}

	p := unsafe.Pointer(v.GObject)
	return C.toGtkPopover(p)
}

func marshalPopover(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapPopover(glib.Take(unsafe.Pointer(c))), nil
}

func wrapPopover(obj *glib.Object) *Popover {
	return &Popover{Bin{Container{Widget{glib.InitiallyUnowned{obj}}}}}
}

//gtk_popover_new()
func PopoverNew(relative IWidget) (*Popover, error) {
	//Takes relative to widget
	var c *C.struct__GtkWidget
	switch {
	case relative == nil:
		c = C.gtk_popover_new(nil)
	default:
		c = C.gtk_popover_new(relative.toWidget())
	}

	if c == nil {
		return nil, nilPtrErr
	}
	return wrapPopover(glib.Take(unsafe.Pointer(c))), nil
}

//gtk_popover_set_default_widget (GtkPopover *popover, GtkWidget *widget);
func (p *Popover) SetDefaultWidget(widget IWidget) {
	C.gtk_popover_set_default_widget(p.native(), widget.toWidget())
}

//GtkWidget *
//gtk_popover_get_default_widget (GtkPopover *popover);
func (p *Popover) GetDefaultWidget() *Widget {
	w := C.gtk_popover_get_default_widget(p.native())
	if w == nil {
		return nil
	}
	return &Widget{glib.InitiallyUnowned{glib.Take(unsafe.Pointer(w))}}
}
