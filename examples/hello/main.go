// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package main

import (
	"fmt"
	"github.com/visualfc/go-iup/iup"
)

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()
	fmt.Println("Hello go-iup!")
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
						iup.Message("About", "GO-IUP\nvisualfc@gmail.com 2011-2012")
					},
				),
			),
		),
	)
	defer dlg.Destroy()
	dlg.Show()
	iup.MainLoop()
}
