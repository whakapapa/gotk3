#include <stdint.h>
#include <stdlib.h>
#include <string.h>

static PangoLayout* toPangoLayout(void *p)
{
	return ( (PangoLayout*) (p) );
}
