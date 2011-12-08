include $(GOROOT)/src/Make.inc

TARG=vfc/iupim
CGOFILES=\
	iupim.go \

CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liupim

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
