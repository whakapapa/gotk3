#include <stdlib.h>

// Type Casting
static GdkAtom toGdkAtom(void *p)
{
	return ((GdkAtom)p);
}

static GdkDevice *toGdkDevice(void *p)
{
	return (GDK_DEVICE(p));
}

static GdkCursor *toGdkCursor(void *p)
{
	return (GDK_CURSOR(p));
}

static GdkDeviceManager *toGdkDeviceManager(void *p)
{
	return (GDK_DEVICE_MANAGER(p));
}

static GdkDisplay *toGdkDisplay(void *p)
{
	return (GDK_DISPLAY(p));
}

static GdkDragContext *toGdkDragContext(void *p)
{
	return (GDK_DRAG_CONTEXT(p));
}

static GdkPixbuf *toGdkPixbuf(void *p)
{
	return (GDK_PIXBUF(p));
}

static gboolean _gdk_pixbuf_save_png(GdkPixbuf *pixbuf, const char *filename, GError ** err, const char *compression)
{
	return gdk_pixbuf_save(pixbuf, filename, "png", err, "compression", compression, NULL);
}

static gboolean _gdk_pixbuf_save_jpeg(GdkPixbuf *pixbuf, const char *filename, GError ** err, const char *quality)
{
	return gdk_pixbuf_save(pixbuf, filename, "jpeg", err, "quality", quality, NULL);
}

static GdkPixbufLoader *toGdkPixbufLoader(void *p)
{
	return (GDK_PIXBUF_LOADER(p));
}

static GdkScreen *toGdkScreen(void *p)
{
	return (GDK_SCREEN(p));
}

static GdkVisual *toGdkVisual(void *p)
{
	return (GDK_VISUAL(p));
}

static GdkWindow *toGdkWindow(void *p)
{
	return (GDK_WINDOW(p));
}

// Type Casting
static GdkMonitor *toGdkMonitor(void *p)
{
	return (GDK_MONITOR(p));
}

static inline gchar** next_gcharptr(gchar** s) { return (s+1); }
