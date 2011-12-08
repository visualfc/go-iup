// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup
/*
#include <stdlib.h>
#include <iup.h>
*/
import "C"
import (
	"strings"
	"unsafe"
)

func Alarm(title, message string, btns ...string) int {
	var b1 *C.char = nil
	var b2 *C.char = nil
	var b3 *C.char = nil
	t := NewCS(title)
	defer FreeCS(t)
	m := NewCS(message)
	defer FreeCS(m)
	l := len(btns)
	if l >= 3 {
		b3 = NewCS(btns[2])
		defer FreeCS(b3)
	}
	if l >= 2 {
		b2 = NewCS(btns[1])
		defer FreeCS(b2)
	}
	if l >= 1 {
		b1 = NewCS(btns[0])
		defer FreeCS(b1)
	}
	return int(C.IupAlarm(t, m, b1, b2, b3))
}

func GetOpenFile(dir string, filter string) (string, bool) {
	dlg := FileDlg(
		Attrs(
			"DIALOGTYPE", "Open",
			"FILTER", filter,
			"DIRECTORY", dir,
			"NOCHANGEDIR", "YES",
			"PARENTDIALOG", GetGlobal("PARENTDIALOG"),
			"ICON", GetGlobal("ICON"),
		),
	)
	dlg.Popup(CENTERPARENT, CENTERPARENT)
	defer dlg.Destroy()
	return dlg.GetAttribute("VALUE"), dlg.GetInt("STATUS") == 0
}

func GetSaveFile(dir string, filter string) (string, bool) {
	dlg := FileDlg(
		Attrs(
			"DIALOGTYPE", "Save",
			"FILTER", filter,
			"DIRECTORY", dir,
			"ALLOWNEW", "YES",
			"NOCHANGEDIR", "YES",
			"PARENTDIALOG", GetGlobal("PARENTDIALOG"),
			"ICON", GetGlobal("ICON"),
		),
	)
	dlg.Popup(CENTERPARENT, CENTERPARENT)
	defer dlg.Destroy()
	return dlg.GetAttribute("VALUE"), dlg.GetInt("STATUS") != -1
}

func GetColor(x, y int, r, g, b *byte) int {
	return int(C.IupGetColor(C.int(x), C.int(y), (*C.uchar)(r), (*C.uchar)(b), (*C.uchar)(b)))
}

/*
The format string must have the following format, notice the "\n" at the end

"text%x[extra]{tip}\n", where:

text is a descriptive text, to be placed to the left of the entry field in a label.

x is the type of the parameter. The valid options are:

b = boolean (shows a True/False toggle, use "int" in C)
i = integer (shows a integer number filtered text box, use "int" in C)
r = real (shows a real number filtered text box, use "float" in C)
a = angle in degrees (shows a real number filtered text box and a dial [if IupControlsOpen is called], use "float" in C)
s = string (shows a text box, use "char*" in C, it must have room enough for your string)
m = multiline string (shows a multiline text box, use "char*" in C, it must have room enough for your string)
l = list (shows a dropdown list box, use "int" in C for the zero based item index selected)
o = list (shows a list of toggles inside a radio, use "int" in C for the zero based item index selected)  (since 3.3)
t = separator (shows a horizontal line separator label, in this case text can be an empty string, not included in parameter count)
f = string (same as s, but also show a button to open a file selection dialog box)
c = string (same as s, but also show a color button to open a color selection dialog box)
n = string (same as s, but also show a font button to open a font selection dialog box) (since 3.3)
u = buttons names (allow to redefine the OK and Cancel button names, and to add a Help button, use [ok,cancel,help] as extra data, can omit one of them, it will use the default name, not included in parameter count) (since 3.1)

bool int float32 float64 string
*/

func GetParam(title string, format string, out ...interface{}) bool {
	extra := strings.Count(format, "%t") + strings.Count(format, "%u")
	count := strings.Count(format, "\n") - extra
	t := NewCS(title)
	defer FreeCS(t)
	f := NewCS(format)
	defer FreeCS(f)
	args := make([]unsafe.Pointer, count+1)
	for i := 0; i < count; i++ {
		args[i] = nil
		switch v := out[i].(type) {
		case *bool:
			p := new(C.int)
			if *v {
				*p = 1
			} else {
				*p = 0
			}
			args[i] = unsafe.Pointer(p)
		case *int:
			args[i] = unsafe.Pointer(v)
		case *uint:
			args[i] = unsafe.Pointer(v)
		case *float32:
			args[i] = unsafe.Pointer(v)
		case *float64:
			args[i] = unsafe.Pointer(v)
		case *string:
			buf := NewCSN(*v, 4096)
			args[i] = unsafe.Pointer(buf)
		default:
			return false
		}
	}
	args[count] = nil
	r := C.IupGetParamv(t, nil, nil, f, C.int(count), C.int(extra), (*unsafe.Pointer)(&args[0]))
	if r == 0 {
		return false
	}
	for i := 0; i < count; i++ {
		switch v := out[i].(type) {
		case *bool:
			*v = *(*int)(args[i]) != 0
		case *string:
			*v = FromCS((*C.char)(args[i]))
		}
	}
	return true
}

func ListDialog(typ int, title string, list []string, opt, max_col, max_lin int, marks []int) int {
	t := NewCS(title)
	defer FreeCS(t)

	clist := NewCSA(list)
	defer FreeCSA(clist)

	r := C.IupListDialog(C.int(typ), t, C.int(len(list)), &clist[0], C.int(opt), C.int(max_col), C.int(max_lin), (*C.int)(unsafe.Pointer(&marks[0])))
	return int(r)
}

func Message(title string, message string) {
	t := NewCS(title)
	defer FreeCS(t)
	m := NewCS(message)
	defer FreeCS(m)
	C.IupMessage(t, m)
}

func LayoutDialog(h *Handle) *Handle {
	return (*Handle)(C.IupLayoutDialog(h.CHandle()))
}

func ElementPropertiesDialog(h *Handle) *Handle {
	return (*Handle)(C.IupElementPropertiesDialog(h.CHandle()))
}

func Image(width, height int, pixels []byte) *Handle {
	return (*Handle)(C.IupImage(C.int(width), C.int(height), (*C.uchar)(&pixels[0])))
}

func ImageRGB(width, height int, pixels []byte) *Handle {
	return (*Handle)(C.IupImageRGB(C.int(width), C.int(height), (*C.uchar)(&pixels[0])))
}

func ImageRGBA(width, height int, pixels []byte) *Handle {
	return (*Handle)(C.IupImageRGBA(C.int(width), C.int(height), (*C.uchar)(&pixels[0])))
}

func GetText(title string, data string) (string, bool) {
	t := NewCS(title)
	defer FreeCS(t)
	d := NewCSN(data, 4096)
	defer FreeCS(d)
	r := C.IupGetText(t, d)
	return FromCS(d), r != 0
}

/*

func GetText(title, value string) (bool, string) {
	return GetTextParam(title, value, "OK", "Cancel")
}
func GetTextParam(title, value, title_ok, title_cancel string) (bool, string) {
	text := Text(
		Attrs("MULTILINE", "YES",
			"EXPAND", "YES",
			"VALUE", value,
			"FONT", "Courier,12",
			"VISIBLELINES", "10",
			"VISIBLECOLUMNS", "50",
		),
	)

	ok := Button(
		Attr("PADDING", "20x5"),
		Attr("TITLE", title_ok),
		func(arg *ButtonAction) {
			arg.Sender.GetDialog().SetAttribute("STATUS", "1")
			arg.Return = CLOSE
		},
	)

	cancel := Button(
		Attr("PADDING", "20x5"),
		Attr("TITLE", title_cancel),
		func(arg *ButtonAction) {
			arg.Sender.GetDialog().SetAttribute("STATUS", "-1")
			arg.Return = CLOSE
		},
	)

	dlg := Dialog(
		"MINBOX=NO,MAXBOX=NO,SIZE=200x150",
		Attr("TITLE", title),
		Attr("ICON", GetGlobal("ICON")),
		Attr("PARENTDIALOG", GetGlobal("PARENTDIALOG")),
		Vbox(
			"MARGIN=10x10",
			"GAP=10",
			text,
			Hbox(
				"MARGIN=0x0",
				"NORMALIZESIZE=HORZONTAL",
				Fill(),
				ok,
				cancel,
			),
		),
	)
	defer dlg.Destroy()
	dlg.SetAttributeHandle("DEFAULTENTER", ok)
	dlg.SetAttributeHandle("DEFAULTESC", cancel)

	dlg.Map()

	text.SetAttribute("VISIBLELINES", "")
	text.SetAttribute("VISIBLECOLUMNS", "")

	dlg.Popup(CENTERPARENT, CENTERPARENT)
	if dlg.GetAttribute("STATUS") == "-1" {
		return false, ""
	}
	return true, text.GetAttribute("VALUE")
}

*/

func Help(url string) {
	curl := NewCS(url)
	defer FreeCS(curl)
	C.IupHelp(curl)
}
