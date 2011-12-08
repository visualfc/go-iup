// hello project main.go
package main

import (
	"fmt"
	"vfc/iup"
)

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()
	fmt.Println("Hello World!")
	mainui()
}

func mainui() {
	dlg := iup.Dialog(
		iup.Attr("TITLE", "Iup Hello World"),
		"RASTERSIZE=400x300",
		iup.Hbox(
			"MARGIN=5x5,GAP=5",
			iup.Vbox(
				iup.Label(
					"EXPAND=HORIZONTAL,ALIGNMENT=ACENTER,FONTSIZE=12",
					iup.Attr("TITLE", "Welcome to GO-IUP"),
				),
				iup.Label(
					"EXPAND=YES,ALIGNMENT=ACENTER:ACENTER,WORDWRAP=YES",
					iup.Attr("TITLE",
						fmt.Sprintf("%s\nVersion: %s\n\n%s\nVersion: %s",
							iup.IupName, iup.IupVersion, iup.Name, iup.Version)),
				),
			),
			iup.Vbox(
				iup.Button(
					"TITLE=OK",
					"SIZE=50x",
					func(arg *iup.ButtonAction) {
						arg.Return = iup.CLOSE
					},
				),
				iup.Button(
					"TITLE=About",
					"SIZE=50x",
					func(arg *iup.ButtonAction) {
						iup.Message("About","GO-IUP\nvisualfc@gmail.com 2011")
					},
				),
			),
		),
	)
	defer dlg.Destroy()
	dlg.Show()
	iup.MainLoop()
}
