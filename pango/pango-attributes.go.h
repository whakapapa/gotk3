#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static PangoColor* toPangoColor(void *p)
{
	return ( (PangoColor*) (p) );
}
