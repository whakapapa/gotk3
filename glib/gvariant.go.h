#ifndef __GVARIANT_GO_H__
#define __GVARIANT_GO_H__

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <glib.h>

// Type Casting

static GVariant *
toGVariant(void *p)
{
	return (GVariant*)p;
}

static GVariantBuilder *
toGVariantBuilder(void *p)
{
	return (GVariantBuilder*)p;
}

static GVariantDict *
toGVariantDict(void *p)
{
	return (GVariantDict*)p;
}

static GVariantIter *
toGVariantIter(void *p)
{
	return (GVariantIter*)p;
}

#endif
