include $(GOROOT)/src/Make.inc

TARG=vfc/iupweb

ifeq ($(GOOS),windows)	
	GOFILES += iupweb_windows.go
else
	CGOFILES +=	iupweb_linux.go
	CGO_CFLAGS = -I../../libs/iup/include
	CGO_LDFLAGS+= -L../../libs/iup -liup -liupweb
endif

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
