#include <stdint.h>
#include <stdlib.h>
#include <string.h>

#include "pango-attributes.go.h"
#include "pango-layout.go.h"

#include "pangocairo.go.h"


static PangoRectangle *
createPangoRectangle(int x, int y, int width, int height)
{
	PangoRectangle *r = (PangoRectangle *)malloc(sizeof(PangoRectangle));
	r->x = x;
	r->y = y;
	r->width = width;
	r->height = height;
	return r;
}
