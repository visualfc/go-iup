include $(GOROOT)/src/Make.inc

TARG=vfc/iupole

CGOFILES=\
	iupole.go \

CGO_CFLAGS= -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liupole

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
