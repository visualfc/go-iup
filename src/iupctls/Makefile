include $(GOROOT)/src/Make.inc

TARG=vfc/iupctls
CGOFILES=\
	iupctls.go \

CGO_CFLAGS = -I../../libs/iup/include
CGO_LDFLAGS+= -L../../libs/iup -liup -liupcontrols

ifeq ($(GOOS),linux)	
CGO_CFLAGS += -I../../libs/cd/include
CGO_LDFLAGS += -L../../libs/cd -liupcd -lcd
endif

#CLEANFILES+=

include $(GOROOT)/src/Make.pkg
