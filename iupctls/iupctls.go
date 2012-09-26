// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupctls

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liupcontrols
#cgo linux CFLAGS : -I../../libs/cd/include
#cgo linux LDFLAGS : -L../../libs/cd -liupcd -lcd
#include <stdlib.h>
#include <iup.h>
#include <iupcontrols.h>
*/
import "C"
import "unsafe"
import "github.com/visualfc/go-iup/iup"

func Open() *iup.Error {
	r := C.IupControlsOpen()
	if r == C.IUP_ERROR {
		return &iup.Error{"IupControlsOpen"}
	}
	return nil
}

func NewCS(s string) *C.char {
	return C.CString(s)
	cs := make([]byte, len(s)+1)
	copy(cs, s)
	return (*C.char)(unsafe.Pointer(&cs[0]))
}

func FreeCS(cs *C.char) {
	C.free(unsafe.Pointer(cs))
}

type IupMatrix struct {
	*iup.Handle
}

func Matrix(a ...interface{}) *IupMatrix {
	return &IupMatrix{iup.Matrix(a...)}
}

func AttachMatrix(h *iup.Handle) *IupMatrix {
	return &IupMatrix{h}
}

func toNative(h iup.IHandle) *C.Ihandle {
	return (*C.Ihandle)(unsafe.Pointer(h.Native()))
}

func (m *IupMatrix) MatSetAttribute(name string, lin, col int, value string) {
	cname := NewCS(name)
	defer FreeCS(cname)
	cvalue := NewCS(value)
	defer FreeCS(cvalue)
	C.IupMatStoreAttribute(toNative(m), cname, C.int(lin), C.int(col), cvalue)
}

func (m *IupMatrix) MatGetAttribute(name string, lin, col int) string {
	cname := NewCS(name)
	defer FreeCS(cname)
	return C.GoString(C.IupMatGetAttribute(toNative(m), cname, C.int(lin), C.int(col)))
}

func (m *IupMatrix) MatSetAttributeData(name string, lin, col int, ptr uintptr) {
	cname := NewCS(name)
	defer FreeCS(cname)
	C.IupMatSetAttribute(toNative(m), cname, C.int(lin), C.int(col), (*C.char)(unsafe.Pointer(ptr)))
}

func (m *IupMatrix) MatGetAttributeData(name string, lin, col int) uintptr {
	cname := NewCS(name)
	defer FreeCS(cname)
	return (uintptr)(unsafe.Pointer(C.IupMatGetAttribute(toNative(m), cname, C.int(lin), C.int(col))))
}

func (m *IupMatrix) MatGetInt(name string, lin, col int) int {
	cname := NewCS(name)
	defer FreeCS(cname)
	return int(C.IupMatGetInt(toNative(m), cname, C.int(lin), C.int(col)))
}

func (m *IupMatrix) MatGetFloat(name string, lin, col int) float32 {
	cname := NewCS(name)
	defer FreeCS(cname)
	return float32(C.IupMatGetFloat(toNative(m), cname, C.int(lin), C.int(col)))
}
