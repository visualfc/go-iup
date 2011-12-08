include $(GOROOT)/src/Make.inc

TARG=vfc/iup

GOFILES = \
	errors.go

CGOFILES = \
	types.go \
	system.go \
	handle.go \
	dialog.go \
	
CGOFILES_linux = \
	callback_linux.go \
	csutil_linux.go
	
CGOFILES_windows = \
	callback_windows.go \
	csutil_windows.go \
	
CGOFILES += $(CGOFILES_$(GOOS))
	
CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg