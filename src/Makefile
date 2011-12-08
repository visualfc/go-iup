all: install clean

install:
	make -C iup         install
	make -C iupctls    install
	make -C iuppplot    install
	make -C iuptuio    install
	make -C iupweb    install
	make -C iupole    install
	make -C iupim    install
	make -C iupimglib    install

clean:
	make -C iup         clean
	make -C iupctls    clean
	make -C iuppplot    clean
	make -C iuptuio    clean
	make -C iupweb    clean
	make -C iupole    clean
	make -C iupim    clean
	make -C iupimglib    clean
	
nuke:
	make -C iup         nuke
	make -C iupctls    nuke
	make -C iuppplot    nuke
	make -C iuptuio    nuke
	make -C iupweb    nuke
	make -C iupole    nuke
	make -C iupim    nuke	
	make -C iupimglib    nuke

format:
	gofmt -w .
