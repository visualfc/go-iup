include $(GOROOT)/src/Make.inc

all: install clean

install:
	make -C iup         install
	make -C iupctls    install
	make -C iuppplot    install
	make -C iuptuio    install
	make -C iupim    install
	make -C iupimglib    install
ifeq ($(GOOS),windows)	
	make -C iupweb    install
	make -C iupole    install
endif

clean:
	make -C iup         clean
	make -C iupctls    clean
	make -C iuppplot    clean
	make -C iuptuio    clean
	make -C iupim    clean
	make -C iupimglib    clean
ifeq ($(GOOS),windows)	
	make -C iupweb    clean
	make -C iupole    clean
endif
	
nuke:
	make -C iup         nuke
	make -C iupctls    nuke
	make -C iuppplot    nuke
	make -C iuptuio    nuke
	make -C iupim    nuke	
	make -C iupimglib    nuke
ifeq ($(GOOS),windows)	
	make -C iupweb    nuke
	make -C iupole    nuke
endif

format:
	gofmt -w .
