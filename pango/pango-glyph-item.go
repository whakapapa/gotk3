package pango

// #include <pango/pango.h>
// #include "pango.go.h"
// #include <stdlib.h>
import "C"
import (
	//	"github.com/andre-hub/gotk3/glib"
	//	"github.com/andre-hub/gotk3/cairo"
	"unsafe"
)

// GlyphItem is a representation of PangoGlyphItem.
type GlyphItem struct {
	pangoGlyphItem *C.PangoGlyphItem
}

// Native returns a pointer to the underlying PangoGlyphItem.
func (v *GlyphItem) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func (v *GlyphItem) native() *C.PangoGlyphItem {
	return (*C.PangoGlyphItem)(unsafe.Pointer(v.pangoGlyphItem))
}
