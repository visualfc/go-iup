// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#include <iup.h>
*/
import "C"
import "unsafe"

func Open() *Error {
	r := C.IupOpen(nil, nil)
	if r == C.IUP_ERROR {
		return &Error{"IupOpen false"}
	}
	return nil
}

func Close() {
	C.IupClose()
}

func MainLoop() int {
	return int(C.IupMainLoop())
}

func LoopStep() int {
	return int(C.IupLoopStep())
}

func LoopStepWait() int {
	return int(C.IupLoopStepWait())
}

func MainLoopLevel() int {
	return int(C.IupMainLoopLevel())
}

func Flush() {
	C.IupFlush()
}

func ExitLoop() {
	C.IupExitLoop()
}

func RecordInput(filename string, mode int) int {
	cname := NewCS(filename)
	defer FreeCS(cname)
	return int(C.IupRecordInput(cname, C.int(mode)))
}

func PlayInput(filename string) int {
	cname := NewCS(filename)
	defer FreeCS(cname)
	return int(C.IupPlayInput(cname))
}

func SetGlobal(name string, value string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	cvalue := NewCS(value)
	defer FreeCS(cvalue)
	C.IupStoreGlobal(cname, cvalue)
}

func GetGlobal(name string) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return FromCS(C.IupGetGlobal(cname))
}

func SetGlobalPtr(name string, value uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupSetGlobal(cname, (*C.char)(unsafe.Pointer(value)))
}

func GetGlobalPtr(name string) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return uintptr(unsafe.Pointer(C.IupGetGlobal(cname)))
}

func ResetGlobal(name string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupResetAttribute(nil, cname)
}

func LoadLEDFile(file string) error {
	cfile := NewCS(file)
	defer FreeCS(cfile)
	err := C.IupLoad(cfile)
	if err != nil {
		return &Error{"IupLoad"}
	}
	return nil
}

func LoadLEDData(data string) error {
	cdata := NewCS(data)
	defer FreeCS(cdata)
	err := C.IupLoadBuffer(cdata)
	if err != nil {
		return &Error{"IupLoadBuffer"}
	}
	return nil
}

func parser(a []interface{}) (children []IHandle, attrs []IAttribute, sattrs []string, funcs []interface{}) {
	for i := 0; i < len(a); i++ {
		switch v := a[i].(type) {
		case IHandle:
			children = append(children, v)
		case IAttribute:
			attrs = append(attrs, v)
		case []IAttribute:
			attrs = append(attrs, v[:]...)
		case string:
			sattrs = append(sattrs, v)
		default:
			funcs = append(funcs, v)
		}
	}
	return
}

func _ParserAttrs(a string) (attrs []IAttribute) {
	var n1, n2 int
	var s1, s2 string
	b1, b2 := true, false
	for n, v := range a {
		if b1 && v == '=' {
			s1 = a[n1:n]
			b2 = true
			n2 = n + 1
		} else if b2 && v == ',' {
			s2 = a[n2:n]
			b1 = true
			n1 = n + 1
			attrs = append(attrs, &Attribute{s1, s2})
		}
	}
	if b1 && b2 {
		attrs = append(attrs, &Attribute{s1, s2})
	}
	return
}

// iup.Attr(key,value), iup.Attrs(k1,v1,v2,v2,...) ,
// string , IHandle , func(arg *Action)
func New(classname string, a ...interface{}) *Handle {
	cname := NewCS(classname)
	defer FreeCS(cname)
	ih := C.IupCreate(cname)
	if ih == nil {
		return nil
	}
	h := (*Handle)(ih)
	if len(a) == 0 {
		return h
	}
	children, attrs, sattrs, funcs := parser(a)
	for i := 0; i < len(children); i++ {
		h.Append(children[i])
	}
	for i := 0; i < len(attrs); i++ {
		h.SetAttribute(attrs[i].Name(), attrs[i].Value())
	}
	for i := 0; i < len(sattrs); i++ {
		h.SetAttributes(sattrs[i])
	}
	for i := 0; i < len(funcs); i++ {
		h.SetCallback(funcs[i])
	}
	return h
}

type ClassInfo struct {
	className   string
	setCallback func(*Handle, interface{}) bool
}

func NewClassInfo(classname string, fn func(*Handle, interface{}) bool) *ClassInfo {
	return &ClassInfo{classname, fn}
}

func (i ClassInfo) ClassName() string {
	return i.className
}

func (i *ClassInfo) SetCallback(h *Handle, fn interface{}) bool {
	return i.setCallback(h, fn)
}

func (i *ClassInfo) GetClassAttributes() []string {
	return GetClassAttributes(i.className)
}

func (i *ClassInfo) GetClassCallbacks() []string {
	return GetClassCallbacks(i.className)
}

func (i *ClassInfo) SetClassDefaultAttribute(name, value string) {
	SetClassDefaultAttribute(i.className, name, value)
}

func GetClassAttributes(classname string) []string {
	c := NewCS(classname)
	defer FreeCS(c)
	total_count := C.IupGetClassAttributes(c, nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetClassAttributes(c, &parray[0], total_count)
	return FromCSA(parray[0:count])
}

func GetClassCallbacks(classname string) []string {
	c := NewCS(classname)
	defer FreeCS(c)
	total_count := C.IupGetClassCallbacks(c, nil, 0)
	if total_count == 0 {
		return nil
	}
	parray := make([]*C.char, total_count)
	count := C.IupGetClassCallbacks(c, &parray[0], total_count)
	return FromCSA(parray[0:count])
}

func SetClassDefaultAttribute(classname, name, value string) {
	c := NewCS(classname)
	defer FreeCS(c)
	n := NewCS(name)
	defer FreeCS(n)
	v := NewCS(value)
	defer FreeCS(v)
	C.IupSetClassDefaultAttribute(c, n, v)
}

var classMap = make(map[string]*ClassInfo)

func RegisterClass(name string, info *ClassInfo) {
	classMap[name] = info
}

func init() {
	RegisterAllClass()
}
