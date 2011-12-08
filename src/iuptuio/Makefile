include $(GOROOT)/src/Make.inc

TARG=vfc/iuptuio
CGOFILES=\
	iuptuio.go \

CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liuptuio

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
