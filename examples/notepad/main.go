package main

import (
	"fmt"
	"time"
	"sync"
	"io/ioutil"
	"vfc/iup"
)

type Notepad struct {
	fileName string
	modify   bool
	dlg      *iup.Handle
	edit     *iup.Handle
	sts1     *iup.Handle
	sts2     *iup.Handle
	timer    *time.Timer
}

var m sync.Mutex

func (pad *Notepad) ShowStatus(h *iup.Handle) {
	if pad.timer != nil {
		pad.timer.Stop()
	}
	tip := h.GetAttribute("TIP")
	pad.sts1.SetAttribute("TITLE", tip)
	pad.timer = time.AfterFunc(1e9, func() {
		pad.sts1.SetAttribute("TITLE", "Ready")
	})
}

func (pad *Notepad) init() *Notepad {
	iup.Menu(
		iup.SubMenu("TITLE=File",
			iup.Menu(
				iup.Item(
					iup.Attr("TITLE", "New\tCtrl+N"),
					iup.Attr("TIP", "New File"),
					func(arg *iup.ItemAction) {
						pad.NewFile()
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
				iup.Item(
					iup.Attr("TITLE", "Open\tCtrl+O"),
					iup.Attr("TIP", "Open File"),
					func(arg *iup.ItemAction) {
						pad.Open()
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
				iup.Item(
					iup.Attr("TITLE", "Save\tCtrl+S"),
					iup.Attr("TIP", "Save File"),
					func(arg *iup.ItemAction) {
						pad.Save()
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
				iup.Item(
					iup.Attr("TITLE", "SaveAs"),
					iup.Attr("TIP", "Save File As..."),
					func(arg *iup.ItemAction) {
						pad.SaveAS()
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
				iup.Separator(),
				iup.Item(
					iup.Attr("TITLE", "Quit"),
					iup.Attr("TIP", "Exit Application"),
					func(arg *iup.ItemAction) {
						pad.CheckModify()
						arg.Return = iup.CLOSE
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
			),
		),
		iup.SubMenu("TITLE=Help",
			iup.Menu(
				iup.Item(
					iup.Attr("TITLE", "About"),
					iup.Attr("TIP", "About Notepad"),
					func(arg *iup.ItemAction) {
						iup.Message("About","\tNotepad 1.0\n\n\tvisualfc@gmail.com 2011\t")
					},
					func(arg *iup.ItemHighlight) {
						pad.ShowStatus(arg.Sender)
					},
				),
			),
		),
	).SetName("main_menu")
	pad.edit = iup.Text(
		"EXPAND=YES",
		"MULTILINE=YES",
		"WORDWRAP=YES",
		"TABSIZE=4",
		"NAME=text",
		func(arg *iup.ValueChanged) {
			pad.SetModify()
		},
		func(arg *iup.TextCaret) {
			pad.sts2.SetAttribute("TITLE", fmt.Sprintf("Lin:%d  Col:%d ", arg.Lin, arg.Col))
		},
		func(arg *iup.CommonKeyAny) {
			key := iup.KeyState(arg.Key)
			if !key.IsCtrl() {
				return
			}
			switch key.Key() {
			case 'N':
				pad.NewFile()
			case 'O':
				pad.Open()
			case 'S':
				pad.Save()
			default:
				return
			}
		},
	)
	pad.sts1 = iup.Label(
		"TITLE=Ready",
		"EXPAND=HORIZONTAL",
		"SIZE=50x",
	)
	pad.sts2 = iup.Label(
		"TITLE=\"Lin:1  Col:1\"",
		"EXPAND=NO",
		"SIZE=60x",
	)
	pad.dlg = iup.Dialog(
		iup.Attrs(
			"MENU", "main_menu",
			"TITLE", "Notepad",
			"SHRINK", "YES",
			"SIZE", "300x200",
		),
		iup.Vbox(
			pad.edit,
			iup.Hbox(
				pad.sts1,
				iup.Fill(),
				pad.sts2,
			),
		),
		func(arg *iup.DialogClose) {
			pad.CheckModify()
			arg.Return = iup.CLOSE
		},
	)
	return pad
}

func (pad *Notepad) CheckModify() {
	if pad.modify == false {
		return
	}
	msg := iup.MessageDlg(
		iup.Attrs(
			"DIALOGTYPE", "WARNING",
			"TITLE", "Notepad",
			"BUTTONS", "YESNO",
			"VALUE", "File is Modify, Save File",
		),
	)
	msg.Popup(iup.CENTERPARENT, iup.CENTERPARENT)
	defer msg.Destroy()
	if msg.GetAttribute("BUTTONRESPONSE") == "2" {
		return
	}
	pad.Save()
}

func (pad *Notepad) Save() {
	fileName := pad.fileName
	if fileName == "" {
		pad.SaveAS()
	} else {
		pad.SaveFile(fileName)
	}
}

func (pad *Notepad) SaveAS() {
	if name, ok := iup.GetSaveFile("", "*.*"); ok {
		pad.SaveFile(name)
	}
}

func (pad *Notepad) NewFile() {
	pad.CheckModify()

	pad.edit.SetAttribute("VALUE", "")
	pad.sts1.SetAttribute("TITLE", "Ready")
	pad.sts2.SetAttribute("TITLE", "Lin:1  Col:1 ")
	pad.SetFileName("")
}

func (pad *Notepad) Open() {
	pad.CheckModify()
	if name, ok := iup.GetOpenFile("", "*.*"); ok {
		pad.OpenFile(name)
	}
}

func (pad *Notepad) OpenFile(fileName string) bool {
	data, e := ioutil.ReadFile(fileName)
	if e != nil {
		return false
	}
	pad.edit.SetAttribute("VALUE", string(data))
	pad.SetFileName(fileName)
	return true
}

func (pad *Notepad) SaveFile(fileName string) bool {
	data := pad.edit.GetAttribute("VALUE")
	e := ioutil.WriteFile(fileName, []byte(data), 0666)
	if e != nil {
		return false
	}
	pad.SetFileName(fileName)
	return true
}

func (pad *Notepad) SetFileName(fileName string) {
	pad.modify = false
	pad.fileName = fileName
	if fileName == "" {
		pad.dlg.SetAttribute("TITLE", "noname - Notepad")
	} else {
		pad.dlg.SetAttribute("TITLE", fileName+" - Notepad")
	}
}

func (pad *Notepad) SetModify() {
	pad.modify = true
	if pad.fileName == "" {
		pad.dlg.SetAttribute("TITLE", "noname* - Notepad")
	} else {
		pad.dlg.SetAttribute("TITLE", pad.fileName+"* - Notepad")
	}
}

func main() {
	e := iup.Open()
	if e != nil {
		fmt.Println(e)
		return
	}
	defer iup.Close()

	pad := new(Notepad).init()
	pad.NewFile()
	pad.dlg.Show()
	iup.MainLoop()
}
