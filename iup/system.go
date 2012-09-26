// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iup

/*
#cgo CFLAGS : -I../../libs/iup/include
#cgo LDFLAGS: -L../../libs/iup -liup
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
