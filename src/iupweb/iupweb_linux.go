// Copyright (C) 2011 visualfc. All rights reserved.
// Use of this source code is governed by a MIT license 
// that can be found in the COPYRIGHT file.

package iupweb

/*
#include <stdlib.h>
#include <iup.h>
#include <iupweb.h>
*/
import "C"
import "vfc/iup"

func Open() *iup.Error {
	r := C.IupWebBrowserOpen()
	if r == C.IUP_ERROR {
		return &iup.Error{"IupWebBrowserOpen"}
	}
	return nil
}
