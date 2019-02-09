// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

// +build !gtk_3_6,!gtk_3_8
// not use this: go build -tags gtk_3_8'. Otherwise, if no build tags are used, GTK 3.10

package gtk

// #include <gtk/gtk.h>
// #include "gtk_since_3_10.go.h"
import "C"

/*
 * GtkWindow
 */
