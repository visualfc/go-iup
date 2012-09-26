// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupweb

import (
	"github.com/visualfc/go-iup/iup"
	"syscall"
)

func Open() *iup.Error {
	h, err := syscall.LoadLibrary("iupweb.dll")
	if err != nil {
		return &iup.Error{"LoadLibrary iupweb.dll false"}
	}
	proc, err := syscall.GetProcAddress(h, "IupWebBrowserOpen")
	if err != nil {
		return &iup.Error{"GetProcAddress IupWebBrowserOpen false"}
	}
	syscall.Syscall(uintptr(proc), 0, 0, 0, 0)
	return nil
}
