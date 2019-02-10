package pango

// #cgo pkg-config: fontconfig gobject-2.0 pango pangocairo
// #include <pango/pango.h>
// #include "pango.go.h"
import "C"

//	"github.com/andre-hub/gotk3/glib"
//	"unsafe"

func init() {

}

/*
* Type conversions
*/

func gbool(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gobool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

/*
* Constantes
*/

const (
	SCALE int = 1024
)
