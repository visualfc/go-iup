// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package main

import (
	"fmt"
	"github.com/visualfc/go-iup/iup"
	"github.com/visualfc/go-iup/iupctls"
	"github.com/visualfc/go-iup/iupim"
	"github.com/visualfc/go-iup/iupimglib"
	"github.com/visualfc/go-iup/iuppplot"
)

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()

	iupctls.Open()
	iuppplot.Open()
	iupimglib.Open()

	ui()
}

func ui() {
	var img *iup.Handle
	tree := iup.Tree(
		"EXPANDALL=YES",
	)
	tree.SetName("tree")
	vbox1 := iup.Vbox(
		"TABTITLE=Standard,MARGIN=0x0,GAP=5",
		iup.Frame("TITLE=Frame,EXPAND=YES",
			iup.Vbox(
				iup.Label("TITLE=\"List Label\""),
				iup.List("EXPAND=YES,1=list1", "2=list2", "3=list3"),
			),
		),
		iup.Hbox(
			iup.List("DROPDOWN=YES,EDITBOX=YES", "1=list1", "2=list2"),
			iup.SpinBox(
				iup.Text("VALUE=100"),
			),
			iup.Toggle("TITLE=Toggle,EXPAND=VERTICAL"),
			//			iup.Val("MIN=1,MAX=100"),
		),
		iup.Frame(
			iup.Radio(
				iup.Hbox(
					iup.Toggle("TITLE=1"),
					iup.Toggle("TITLE=2"),
					iup.Toggle("TITLE=3"),
					iup.Toggle("TITLE=4"),
				),
			),
		),
		iup.Text("EXPAND=HORIZONTAL"),
		iup.ProgressBar("MIN=0,MAX=100,VALUE=50,EXPAND=HORIZONTAL"),
		tree,
	)
	vbox2 := iup.Vbox(
		"TABTITLE=Text",
		iup.Text("EXPAND=YES,MULTILINE=YES"),
	)
	vbox3 := iup.Vbox(
		"TABTITLE=Colorbar",
		iup.Colorbar(
			"ORIENTATION=HORIZONTAL",
			"NUM_PARTS=2",
			"SHOW_SECONDARY=YES",
			func(arg *iup.ColorbarSelect) {
				fmt.Println(arg)
			},
		),
	)
	vbox4 := iup.Vbox(
		"TABTITLE=ColorBrowser",
		iup.ColorBrowser(
			"EXPAND=YES",
			func(arg *iup.ColorBrowserDrag) {
				fmt.Printf("R=%d,G=%d,B=%d\n", arg.R, arg.G, arg.B)
			},
		),
	)
	vbox5 := iup.Vbox(
		"TABTITLE=Dial",
		iup.Fill(),
		iup.Dial(),
		iup.Dial("ORIENTATION=VERTICAL"),
		iup.Fill(),
	)
	pplot := iuppplot.PPlot(
		iup.Attr("TITLE", "A simple XY Plot"),
		iup.Attrs(
			"MARGINBOTTOM", "35",
			"MARGINLEFT", "35",
			"AXS_XLABEL", "X",
			"AXS_YLABEL", "Y",
		),
	)
	pplot.Begin(0)
	pplot.Add(0, 0)
	pplot.Add(5, 5)
	pplot.Add(10, 7)
	pplot.End()
	vbox6 := iup.Vbox(
		"TABTITLE=PPlot",
		pplot,
	)
	vbox7 := iup.Vbox(
		"TABTITLE=Matrix",
		iup.Matrix(
			"NUMCOL=5,NUMLIN=10,NUMCOL_VISIBLE=5,NUMLIN_VISIBLE=3,WIDTHDEF=34",
			"RESIZEMATRIX=YES",
		),
	)
	vbox8 := iup.Vbox(
		"TABTITLE=Image",
		iup.Hbox(
			"GAP=10",
			iup.Button("IMAGE=IUP_FileNew"),
			iup.Button("IMAGE=IUP_FileOpen"),
			iup.Button(
				"TITLE=SelectImg",
				func(arg *iup.ButtonAction) {
					file, ok := iup.GetOpenFile("", "*.png;*.jpg;*.bmp;*.jpeg")
					if ok {
						label := arg.Sender.GetDialogChild("img_label")
						if img != nil {
							img.Destroy()
						}
						img = iupim.LoadImage(file)
						if img != nil {
							img.SetName("img_label_image")
							label.(*iup.Handle).SetAttribute("IMAGE", "img_label_image")
							label.(*iup.Handle).Refresh()
						}
					}
				},
			),
		),
		iup.Label("IMAGE=img_label_image", "NAME=img_label"),
	)
	dlg := iup.Dialog(
		iup.Attr("TITLE", "GO-IUP Demo 1.0"),
		"SIZE=350x200",
		"SHRINK=YES",
		"MARGIN=10x10",
		iup.Tabs(
			vbox1,
			vbox2,
			vbox3,
			vbox4,
			vbox5,
			vbox6,
			vbox7,
			vbox8,
		),
	)
	dlg.Show()

	tree.SetAttribute("ADDBRANCH", "Item2")
	tree.SetAttribute("ADDLEAF1", "leaf3")
	tree.SetAttribute("ADDLEAF2", "leaf4")
	tree.SetAttribute("ADDBRANCH", "Item1")
	tree.SetAttribute("ADDLEAF1", "leaf1")
	tree.SetAttribute("ADDLEAF2", "leaf2")

	defer dlg.Destroy()
	iup.MainLoop()
}
