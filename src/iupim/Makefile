include $(GOROOT)/src/Make.inc

TARG=vfc/iupim
CGOFILES=\
	iupim.go \

CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liupim

ifeq ($(GOOS),linux)	
CGO_CFLAGS += -I../../libs/im/include
CGO_LDFLAGS += -L../../libs/im -lim
endif

include $(GOROOT)/src/Make.pkg
