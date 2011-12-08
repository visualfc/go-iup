// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupweb

import (
	"syscall"
	"vfc/iup"
)

func Open() *iup.Error {
	h, err := syscall.LoadLibrary("iupweb.dll")
	if err != 0 {
		return &iup.Error{"LoadLibrary iupweb.dll false"}
	}
	proc, err := syscall.GetProcAddress(h, "IupWebBrowserOpen")
	if err != 0 {
		return &iup.Error{"GetProcAddress IupWebBrowserOpen false"}
	}
	syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	return nil
}
