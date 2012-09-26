// Copyright (C) 2011-2012 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupweb

/*
#cgo linux CFLAGS: -I../../libs/iup/include
#cgo linux LDFLAGS: -L../../libs/iup -liup -liupweb

#include <stdlib.h>
#include <iup.h>
#include <iupweb.h>
*/
import "C"
import "github.com/visualfc/go-iup/iup"

func Open() *iup.Error {
	r := C.IupWebBrowserOpen()
	if r == C.IUP_ERROR {
		return &iup.Error{"IupWebBrowserOpen"}
	}
	return nil
}
