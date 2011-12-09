include $(GOROOT)/src/Make.inc

TARG=vfc/iupimglib
CGOFILES=\
	iupimglib.go \

CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liupimglib

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
